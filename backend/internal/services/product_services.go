package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
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
		return errors.New("PRODUCT_NAME_REQUIRED")
	}
	if productCreationDTO.Price < 0 {
		return errors.New("PRODUCT_PRICE_INVALID")
	}
	if productCreationDTO.StockQuantity < 0 {
		return errors.New("PRODUCT_QUANTITY_INVALID")
	}
	if productCreationDTO.Description == "" {
		return errors.New("PRODUCT_DESCRIPTION_REQUIRED")
	}
	if err := service.repository.InsertNewProduct(ctx, productCreationDTO); err != nil {
		return err
	}
	return nil
}

func (service *ProductService) RemoveProduct(ctx context.Context, productId int) error {
	if productId < 0 {
		return errors.New("INVALID_PRODUCT_ID")
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
			return fmt.Errorf("INVALID_PRICE")
		}
		updateFields["price"] = *dto.Price
	}
	if dto.StockQuantity != nil {
		if *dto.StockQuantity < 0 {
			return fmt.Errorf("INVALID_STOCKQUANTITY")
		}
		updateFields["stock_quantity"] = *dto.StockQuantity
	}

	if len(updateFields) == 0 {
		return nil
	}

	return service.repository.UpdateProduct(ctx, productId, updateFields)
}
