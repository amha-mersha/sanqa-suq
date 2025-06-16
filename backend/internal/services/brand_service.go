package services

import (
	"context"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type BrandService struct {
	brandRepo *repositories.BrandRepository
}

func NewBrandService(brandRepo *repositories.BrandRepository) *BrandService {
	return &BrandService{
		brandRepo: brandRepo,
	}
}

func (s *BrandService) GetBrands(ctx context.Context) ([]*models.Brands, error) {
	return s.brandRepo.FetchAllBrands(ctx)
}

func (s *BrandService) GetBrandByID(ctx context.Context, id int) (*models.Brands, error) {
	return s.brandRepo.FetchBrandByID(ctx, id)
}

func (s *BrandService) CreateBrand(ctx context.Context, brand *dtos.CreateBrandRequest) (*models.Brands, error) {
	newBrand := &models.Brands{
		Name:        brand.Name,
		Description: brand.Description,
	}
	return s.brandRepo.InsertBrand(ctx, newBrand)
}
