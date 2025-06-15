package services

import (
	"context"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type ProductService struct {
	repository *repositories.ProductRepository
}

func NewProductService(repository *repositories.ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) AddNewProduct(ctx context.Context, productCreationDTO *dtos.CreateProductDTO) error {
	if productCreationDTO.Name == "" {
		return errs.BadRequest("PRODUCT_NAME_REQUIRED", nil)
	}
	if productCreationDTO.Price < 0 {
		return errs.BadRequest("PRODUCT_PRICE_INVALID", nil)
	}
	if productCreationDTO.StockQuantity < 0 {
		return errs.BadRequest("PRODUCT_STOCK_QUANTITY_INVALID", nil)
	}
	if productCreationDTO.Description == "" {
		return errs.BadRequest("PRODUCT_DESCRIPTION_REQUIRED", nil)
	}
	if err := service.repository.InsertNewProduct(ctx, productCreationDTO); err != nil {
		return err
	}
	return nil
}

func (service *ProductService) RemoveProduct(ctx context.Context, productId int) error {
	if productId < 0 {
		return errs.BadRequest("INVALID_PRODUCT_ID", nil)
	}
	return service.RemoveProduct(ctx, productId)
}

func (service *ProductService) UpdateProduct(ctx context.Context, productId int, dto *dtos.ProductUpdateDTO) error {
	updateFields := make(map[string]interface{})

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
