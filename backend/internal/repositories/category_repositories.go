package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal"
	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v5"
)

type CategoryRepository struct {
	DB *database.DB
}

func NewCategoryRepository(db *database.DB) *CategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (repositories *CategoryRepository) GetAllCategories(ctx context.Context) ([]models.Categories, error) {
	query := `SELECT category_id, category_name, parent_category_id FROM categories;`
	rows, err := repositories.DB.Pool.Query(ctx, query)
	if err != nil {
		return nil, internal.InternalError("Failed to fetch categories", err)
	}
	defer rows.Close()

	categories, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Categories, error) {
		var c models.Categories
		err := row.Scan(&c.CategoryID, &c.Name, &c.ParentCategoryID)
		return c, err
	})
	if err != nil {
		return nil, internal.InternalError("Failed to fetch categories", err)
	}
	return categories, nil
}

func (r *CategoryRepository) InsertCategory(ctx context.Context, newCategory *models.Categories) error {
	query := `INSERT INTO categories (category_name, parent_category_id) 
	          VALUES ($1, $2) RETURNING category_id;`
	err := r.DB.Pool.QueryRow(ctx, query, newCategory.Name, newCategory.ParentCategoryID).
		Scan(&newCategory.CategoryID)
	if err != nil {
		return internal.InternalError("Failed to insert category", err)
	}
	return nil
}

func (r *CategoryRepository) GetCategoryById(ctx context.Context, categoryId string) (*models.Categories, error) {
	query := `SELECT category_id, category_name, parent_category_id 
	          FROM categories WHERE category_id = $1;`
	row := r.DB.Pool.QueryRow(ctx, query, categoryId)
	var category models.Categories
	err := row.Scan(&category.CategoryID, &category.Name, &category.ParentCategoryID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, internal.NotFound("Category not found", nil)
		}
		return nil, internal.InternalError("Failed to fetch category", err)
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, categoryId int, fields map[string]any) error {
	setClause := []string{}
	args := []any{}
	i := 1
	for key, value := range fields {
		setClause = append(setClause, fmt.Sprintf("%s = $%d", key, i))
		args = append(args, value)
		i++
	}
	args = append(args, fields["category_id"])
	query := fmt.Sprintf("UPDATE categories SET %s WHERE category_id = $%d", strings.Join(setClause, ", "), i)
	_, err := r.DB.Pool.Exec(ctx, query, args...)
	if err != nil {
		return internal.InternalError("Failed to update category", err)
	}
	return nil
}
func (r *CategoryRepository) DeleteCategory(ctx context.Context, categoryId int) error {
	query := `DELETE FROM categories WHERE category_id = $1;`
	result, err := r.DB.Pool.Exec(ctx, query, categoryId)
	if err != nil {
		return internal.InternalError("Failed to delete category", err)
	}
	if result.RowsAffected() == 0 {
		return internal.NotFound("Category not found", nil)
	}
	return nil
}
