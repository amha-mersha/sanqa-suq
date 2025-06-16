package services

import (
	"context"
	"fmt"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) AddNewProduct(ctx context.Context, productCreationDTO *dtos.CreateProductDTO) (*models.Products, error) {
	if productCreationDTO.Name == "" {
		return nil, errs.BadRequest("PRODUCT_NAME_REQUIRED", nil)
	}
	if productCreationDTO.Price < 0 {
		return nil, errs.BadRequest("PRODUCT_PRICE_INVALID", nil)
	}
	if productCreationDTO.StockQuantity < 0 {
		return nil, errs.BadRequest("PRODUCT_STOCK_QUANTITY_INVALID", nil)
	}
	if productCreationDTO.Description == "" {
		return nil, errs.BadRequest("PRODUCT_DESCRIPTION_REQUIRED", nil)
	}
	return service.repository.InsertNewProduct(ctx, productCreationDTO)
}

func (service *ProductService) RemoveProduct(ctx context.Context, productId int) error {
	if productId < 0 {
		return errs.BadRequest("INVALID_PRODUCT_ID", nil)
	}
	return service.RemoveProduct(ctx, productId)
}

func (service *ProductService) UpdateProduct(ctx context.Context, productId int, dto *dtos.ProductUpdateDTO) error {
	updateFields := make(map[string]any)

	if dto.Name != nil {
		updateFields["name"] = *dto.Name
	}
	if dto.Description != nil {
		updateFields["description"] = *dto.Description
	}
	if dto.BrandID != nil {
		updateFields["brand_id"] = *dto.BrandID
	}
	if dto.CategoryID != nil {
		updateFields["category_id"] = *dto.CategoryID
	}
	if dto.Price != nil {
		if *dto.Price < 0 {
			return errs.BadRequest("INVALID_PRICE", nil)
		}
		updateFields["price"] = *dto.Price
	}
	if dto.StockQuantity != nil {
		if *dto.StockQuantity < 0 {
			return errs.BadRequest("INVALID_STOCK_QUANTITY", nil)
		}
		updateFields["stock_quantity"] = *dto.StockQuantity
	}

	if len(updateFields) == 0 {
		return nil
	}

	return service.repository.UpdateProduct(ctx, productId, updateFields)
}

func (service *ProductService) GetProduct(ctx context.Context, productId int) (*models.Products, error) {
	if productId < 0 {
		return nil, errs.BadRequest("INVALID_PRODUCT_ID", nil)
	}
	return service.repository.FindProductByID(ctx, productId)
}

func (service *ProductService) GetAllProducts(ctx context.Context) ([]*models.Products, error) {
	products, err := service.repository.FetchAllProducts(ctx)
	if err != nil {
		return nil, errs.InternalError("failed to fetch all products", err)
	}
	return products, nil
}

func (service *ProductService) GetProductsByCategoryID(ctx context.Context, categoryId int) ([]*models.Products, error) {
	if categoryId <= 0 {
		return nil, errs.BadRequest("invalid category ID", fmt.Errorf("categoryId must be positive, got %d", categoryId))
	}
	category, err := service.repository.FindCategoryByID(ctx, categoryId)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errs.NotFound("CATEGORY_NOT_FOUND", nil)
	}
	return service.repository.FetchProductsByCategoryID(ctx, categoryId)
}
