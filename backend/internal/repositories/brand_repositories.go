package repositories

import (
	"context"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgx/v5"
)

type BrandRepository struct {
	DB *database.DB
}

func NewBrandRepository(db *database.DB) *BrandRepository {
	return &BrandRepository{
		DB: db,
	}
}

func (r *BrandRepository) FetchAllBrands(ctx context.Context) ([]*models.Brands, error) {
	query := "SELECT brand_id, name, description FROM brands"
	rows, err := r.DB.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brands []*models.Brands
	for rows.Next() {
		brand := &models.Brands{}
		if err := rows.Scan(&brand.BrandID, &brand.Name, &brand.Description); err != nil {
			return nil, errs.InternalError("failed to scan brand row", err)
		}
		brands = append(brands, brand)
	}

	if err = rows.Err(); err != nil {
		return nil, errs.InternalError("error iterating brand rows", err)
	}

	return brands, nil
}

func (r *BrandRepository) FetchBrandByID(ctx context.Context, id int) (*models.Brands, error) {
	query := "SELECT * FROM brands WHERE brand_id = $1"
	row := r.DB.Pool.QueryRow(ctx, query, id)
	var brand models.Brands
	err := row.Scan(&brand.BrandID, &brand.Name, &brand.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NotFound("brand not found", err)
		}
		return nil, errs.InternalError("failed to fetch brand by ID", err)
	}
	return &brand, nil
}

func (r *BrandRepository) InsertBrand(ctx context.Context, brand *models.Brands) (*models.Brands, error) {
	query := "INSERT INTO brands (name, description) VALUES ($1, $2) RETURNING brand_id"
	err := r.DB.Pool.QueryRow(ctx, query, brand.Name, brand.Description).Scan(&brand.BrandID)
	if err != nil {
		return nil, errs.InternalError("failed to insert brand", err)
	}
	return brand, nil
}
