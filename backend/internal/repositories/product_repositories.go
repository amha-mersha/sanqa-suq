package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	"github.com/amha-mersha/sanqa-suq/internal/models"
)

type ProductRepository struct {
	DB *database.DB
}

func NewProductRepository(db *database.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repository *ProductRepository) FindCategoryByID(ctx context.Context, categoryId int) (*models.Categories, error) {
	var category models.Categories
	query := `SELECT category_id,name, parent_category_id  FROM	categories c WHERE c.category_id = $1`
	err := repository.DB.Pool.QueryRow(ctx, query, categoryId).Scan(&category.CategoryID, &category.Name, &category.ParentCategoryID)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository *ProductRepository) FindBrandByID(ctx context.Context, brandId int) (*models.Brands, error) {
	var brand models.Brands
	query := `SELECT brand_id,name, description FROM brands b WHERE b.brand_id = $1`
	err := repository.DB.Pool.QueryRow(ctx, query, brandId).Scan(&brand.BrandID, &brand.Name, &brand.Description)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (repository *ProductRepository) FindProductByID(ctx context.Context, id int) (*models.Products, error) {
	query := `SELECT product_id, category_id, brand_id, name, description, price, stock_quantity FROM products WHERE product_id = $1`
	var product models.Products
	err := repository.DB.Pool.QueryRow(ctx, query, id).Scan(&product.ProductID, &product.CategoryID, &product.BrandID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (repository *ProductRepository) InsertNewProduct(ctx context.Context, product *dtos.CreateProductDTO) error {
	query := `INSERT INTO products (category_id, brand_id, name, description, price, stock_quantity) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repository.DB.Pool.Exec(ctx, query, product.CategoryID, product.BrandID, product.Name, product.Description, product.Price, product.StockQuantity)
	return err
}

func (repository *ProductRepository) DeleteProductByID(ctx context.Context, productID int) error {
	query := `DELETE FROM products WHERE products.product_id = $1`
	result, err := repository.DB.Pool.Exec(ctx, query, productID)
	if err != nil {
		if result.RowsAffected() == 0 {
			return errors.New("PRODUCT_NOT_FOUND")
		}
	}
	return err
}

func (repository *ProductRepository) UpdateProduct(ctx context.Context, productId int, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	setClauses := []string{}
	args := []interface{}{}
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
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("PRODUCT_WITH_ID_%d_NOT_FOUND", productId)
	}

	return nil
}
