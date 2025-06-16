package services

import (
	"context"
	"time"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
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

func (s *BuildService) CreateBuild(ctx context.Context, userID string, req *dtos.CreateBuildRequestDTO) (*dtos.BuildResponseDTO, error) {
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
		Items: make([]struct {
			ProductID    int     `json:"product_id"`
			Quantity     int     `json:"quantity"`
			ProductName  string  `json:"product_name"`
			Price        float64 `json:"price"`
			Description  string  `json:"description"`
			BrandName    string  `json:"brand_name"`
			CategoryName string  `json:"category_name"`
		}, len(result.Items)),
	}

	for i, item := range result.Items {
		response.Items[i] = struct {
			ProductID    int     `json:"product_id"`
			Quantity     int     `json:"quantity"`
			ProductName  string  `json:"product_name"`
			Price        float64 `json:"price"`
			Description  string  `json:"description"`
			BrandName    string  `json:"brand_name"`
			CategoryName string  `json:"category_name"`
		}{
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
	return s.buildRepo.GetBuildByID(ctx, buildID)
}

func (s *BuildService) UpdateBuild(ctx context.Context, buildID string, userID string, req *dtos.UpdateBuildRequestDTO) (*dtos.BuildResponseDTO, error) {
	// Convert DTO items to model items
	var items []models.BuildItem
	if req.Items != nil {
		items = make([]models.BuildItem, len(req.Items))
		for i, item := range req.Items {
			items[i] = models.BuildItem{
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			}
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
		Items: make([]struct {
			ProductID    int     `json:"product_id"`
			Quantity     int     `json:"quantity"`
			ProductName  string  `json:"product_name"`
			Price        float64 `json:"price"`
			Description  string  `json:"description"`
			BrandName    string  `json:"brand_name"`
			CategoryName string  `json:"category_name"`
		}, len(build.Items)),
	}

	for i, item := range build.Items {
		response.Items[i] = struct {
			ProductID    int     `json:"product_id"`
			Quantity     int     `json:"quantity"`
			ProductName  string  `json:"product_name"`
			Price        float64 `json:"price"`
			Description  string  `json:"description"`
			BrandName    string  `json:"brand_name"`
			CategoryName string  `json:"category_name"`
		}{
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
