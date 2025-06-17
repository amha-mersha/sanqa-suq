package services

import (
	"context"
	"time"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type BuildService struct {
	buildRepo *repositories.BuildRepository
}

func NewBuildService(buildRepo *repositories.BuildRepository) *BuildService {
	return &BuildService{
		buildRepo: buildRepo,
	}
}

func (s *BuildService) validateBuildItems(items []dtos.BuildItemDTO) error {
	if len(items) == 0 {
		return errs.BadRequest("build must have at least one item", nil)
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return errs.BadRequest("quantity must be greater than 0", nil)
		}
	}

	return nil
}

func (s *BuildService) CreateBuild(ctx context.Context, userID string, req *dtos.CreateBuildRequestDTO) (*dtos.BuildResponseDTO, error) {
	if err := s.validateBuildItems(req.Items); err != nil {
		return nil, err
	}

	// Convert DTO to model
	build := &models.CustomBuild{
		UserID: userID,
		Name:   req.Name,
	}

	items := make([]models.BuildItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.BuildItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	// Create build in repository
	result, err := s.buildRepo.CreateBuild(ctx, build, items)
	if err != nil {
		return nil, err
	}

	// Convert model to response DTO
	response := &dtos.BuildResponseDTO{
		BuildID:    result.BuildID,
		UserID:     result.UserID,
		Name:       result.Name,
		CreatedAt:  result.CreatedAt.Format(time.RFC3339),
		TotalPrice: result.TotalPrice,
		Items:      make([]dtos.BuildItemResponseDTO, len(result.Items)),
	}

	for i, item := range result.Items {
		response.Items[i] = dtos.BuildItemResponseDTO{
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			ProductName:  item.ProductName,
			Price:        item.Price,
			Description:  item.Description,
			BrandName:    item.BrandName,
			CategoryName: item.CategoryName,
		}
	}

	return response, nil
}

func (s *BuildService) GetUserBuilds(ctx context.Context, userID string) ([]models.BuildWithItems, error) {
	return s.buildRepo.GetUserBuilds(ctx, userID)
}

func (s *BuildService) GetBuildByID(ctx context.Context, buildID string) (*models.BuildWithItems, error) {
	if buildID == "" {
		return nil, errs.BadRequest("build ID is required", nil)
	}
	return s.buildRepo.GetBuildByID(ctx, buildID)
}

func (s *BuildService) UpdateBuild(ctx context.Context, buildID string, userID string, req *dtos.UpdateBuildRequestDTO) (*dtos.BuildResponseDTO, error) {
	if buildID == "" {
		return nil, errs.BadRequest("build ID is required", nil)
	}

	if err := s.validateBuildItems(req.Items); err != nil {
		return nil, err
	}

	// Convert DTO items to model items
	items := make([]models.BuildItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.BuildItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
	}

	// Update the build
	build, err := s.buildRepo.UpdateBuild(ctx, buildID, userID, &req.Name, items)
	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	response := &dtos.BuildResponseDTO{
		BuildID:    build.BuildID,
		UserID:     build.UserID,
		Name:       build.Name,
		CreatedAt:  build.CreatedAt.Format(time.RFC3339),
		TotalPrice: build.TotalPrice,
		Items:      make([]dtos.BuildItemResponseDTO, len(build.Items)),
	}

	for i, item := range build.Items {
		response.Items[i] = dtos.BuildItemResponseDTO{
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			ProductName:  item.ProductName,
			Price:        item.Price,
			Description:  item.Description,
			BrandName:    item.BrandName,
			CategoryName: item.CategoryName,
		}
	}

	return response, nil
}

func (s *BuildService) GetCompatibleProducts(ctx context.Context, categoryID int, selectedItems []int) ([]dtos.CompatibleProductDTO, error) {
	products, err := s.buildRepo.GetCompatibleProducts(ctx, categoryID, selectedItems)
	if err != nil {
		return nil, err
	}

	// Convert model to DTO
	response := make([]dtos.CompatibleProductDTO, len(products))
	for i, product := range products {
		response[i] = dtos.CompatibleProductDTO{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			Price:        product.Price,
			Description:  product.Description,
			BrandName:    product.BrandName,
			CategoryName: product.CategoryName,
			Specs:        product.Specs,
		}
	}

	return response, nil
}
