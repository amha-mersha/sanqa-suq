package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type ProductRepository struct {
	DB *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repository *ProductRepository) FetchAllProducts(ctx context.Context) ([]*models.Products, error) {
	query := `SELECT product_id, category_id, brand_id, name, description, price, stock_quantity FROM products`
	rows, err := repository.DB.Pool.Query(ctx, query)
	if err != nil {
		return nil, errs.InternalError("failed to fetch products", err)
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[*models.Products])
	if err != nil {
		return nil, errs.InternalError("failed to collect product rows", err)
	}
	return products, nil
}

func (repository *ProductRepository) FindCategoryByID(ctx context.Context, categoryId int) (*models.Categories, error) {
	var category models.Categories
	query := `SELECT category_id,name, parent_category_id  FROM	categories c WHERE c.category_id = $1`
	err := repository.DB.Pool.QueryRow(ctx, query, categoryId).Scan(&category.CategoryID, &category.Name, &category.ParentCategoryID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("category with id %d not found", categoryId), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to find category with id %d", categoryId), err)
	}
	return &category, nil
}

func (repository *ProductRepository) FetchProductsByCategoryID(ctx context.Context, categoryId int) ([]*models.Products, error) {
	if categoryId <= 0 {
		return nil, errs.BadRequest("invalid category ID", fmt.Errorf("categoryId must be positive, got %d", categoryId))
	}
	query := `
		WITH RECURSIVE category_hierarchy AS (
			SELECT category_id
			FROM categories
			WHERE category_id = $1
			UNION ALL
			SELECT c.category_id
			FROM categories c
			INNER JOIN category_hierarchy ch ON c.parent_category_id = ch.category_id
		)
		SELECT p.product_id, p.category_id, p.brand_id, p.name, p.description, p.price, p.stock_quantity
		FROM products p
		INNER JOIN category_hierarchy ch ON p.category_id = ch.category_id
	`
	rows, err := repository.DB.Pool.Query(ctx, query, categoryId)
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to fetch products for category ID %d and its descendants", categoryId), err)
	}
	defer rows.Close()

	products, err := pgx.CollectRows(rows, pgx.RowToStructByPos[*models.Products])
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to scan products for category ID %d and its descendants", categoryId), err)
	}

	if len(products) == 0 {
		return nil, errs.NotFound(fmt.Sprintf("no products found for category ID %d or its descendants", categoryId), nil)
	}

	return products, nil
}

func (repository *ProductRepository) FindBrandByID(ctx context.Context, brandId int) (*models.Brands, error) {
	var brand models.Brands
	query := `SELECT brand_id,name, description FROM brands b WHERE b.brand_id = $1`
	err := repository.DB.Pool.QueryRow(ctx, query, brandId).Scan(&brand.BrandID, &brand.Name, &brand.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("brand with id %d not found", brandId), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to find brand with id %d", brandId), err)
	}
	return &brand, nil
}

func (repository *ProductRepository) FindProductByID(ctx context.Context, id int) (*models.Products, error) {
	query := `SELECT product_id, category_id, brand_id, name, description, price, stock_quantity FROM products WHERE product_id = $1`
	var product models.Products
	err := repository.DB.Pool.QueryRow(ctx, query, id).Scan(&product.ProductID, &product.CategoryID, &product.BrandID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("product with id %d not found", id), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to find product with id %d", id), err)
	}
	return &product, nil
}

func (repository *ProductRepository) InsertNewProduct(ctx context.Context, productDTO *dtos.CreateProductDTO) (*models.Products, error) {
	query := `
		INSERT INTO products (
			category_id,
			brand_id,
			name,
			description,
			price,
			stock_quantity
		) VALUES (
			$1, $2, $3, $4, $5, $6
		)
		RETURNING product_id, category_id, brand_id, name, description, price, stock_quantity, created_at;
	`

	newProduct := &models.Products{}

	err := repository.DB.Pool.QueryRow(
		ctx,
		query,
		productDTO.CategoryID,
		productDTO.BrandID,
		productDTO.Name,
		productDTO.Description,
		productDTO.Price,
		productDTO.StockQuantity,
	).Scan(
		&newProduct.ProductID,
		&newProduct.CategoryID,
		&newProduct.BrandID,
		&newProduct.Name,
		&newProduct.Description,
		&newProduct.Price,
		&newProduct.StockQuantity,
		&newProduct.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23503" {
				return nil, errs.BadRequest("FOREIGN_KEY_VIOLATION", fmt.Errorf("invalid category_id or brand_id: %w", err))
			}
		}
		if err == pgx.ErrNoRows {
			return nil, errs.InternalError("failed to insert new product: no rows returned", err)
		}
		return nil, errs.InternalError("failed to insert new product", err)
	}

	return newProduct, nil
}

func (repository *ProductRepository) DeleteProductByID(ctx context.Context, productID int) error {
	query := `DELETE FROM products WHERE products.product_id = $1`
	result, err := repository.DB.Pool.Exec(ctx, query, productID)
	if err != nil {
		if result.RowsAffected() == 0 {
			return errs.NotFound(fmt.Sprintf("product with id %d not found", productID), err)
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return errs.NotFound(fmt.Sprintf("product with id %d not found", productID), err)
		}
		return errs.InternalError(fmt.Sprintf("failed to delete product with id %d", productID), err)
	}
	return nil
}

func (repository *ProductRepository) UpdateProduct(ctx context.Context, productId int, fields map[string]any) error {
	if len(fields) == 0 {
		return nil
	}

	setClauses := []string{}
	args := []any{}
	i := 1

	for col, val := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}

	args = append(args, productId)
	query := fmt.Sprintf("UPDATE products SET %s WHERE id = $%d", strings.Join(setClauses, ", "), i)

	cmdTag, err := repository.DB.Pool.Exec(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errs.NotFound(fmt.Sprintf("product with id %d not found", productId), err)
		}
		return errs.InternalError(fmt.Sprintf("failed to update product with id %d", productId), err)
	}
	if cmdTag.RowsAffected() == 0 {
		return errs.NotFound(fmt.Sprintf("product with id %d not found", productId), nil)
	}

	return nil
}
