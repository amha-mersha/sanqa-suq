package services

import (
	"context"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type CategoryService struct {
	repository *repositories.CategoryRepository
}

func NewCategoryService(repository *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}
func (service *CategoryService) GetCategoryList(ctx context.Context) ([]models.Categories, error) {
	return service.repository.GetAllCategories(ctx)
}

func (serivce *CategoryService) CreateCategory(ctx context.Context, newCategory *dtos.CreateCategoryDTO) (*models.Categories, error) {
	category := models.Categories{
		Name:             newCategory.Name,
		ParentCategoryID: newCategory.ParentCategoryID,
	}
	return serivce.repository.InsertCategory(ctx, &category)
}

func (s *CategoryService) GetCategoryById(ctx context.Context, categoryId string) (*models.Categories, error) {
	category, err := s.repository.GetCategoryById(ctx, categoryId)
	return category, err
}

func (s *CategoryService) GetCategoryWithChildren(ctx context.Context, categoryId string, limit int) (map[string]any, error) {
	return s.repository.FetchCategoryTree(ctx, categoryId, limit)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, categoryId string, updateData *dtos.UpdateCategoryDTO) error {
	updateFields := make(map[string]any)
	convertedCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		return errs.BadRequest("INVALID_CATEGORY_ID", err)
	}
	if updateData.Name != nil {
		if len(*updateData.Name) == 0 {
			return errs.BadRequest("CATEGORY_NAME_CANNOT_BE_EMPTY", nil)
		}
		updateFields["category_name"] = *updateData.Name
	}
	if updateData.ParentCategoryID != nil {
		updateFields["parent_category_id"] = *updateData.ParentCategoryID
	}
	if len(updateFields) == 0 {
		return errs.BadRequest("NO_FIELDS_TO_UPDATE", nil)
	}
	return s.repository.UpdateCategory(ctx, convertedCategoryId, updateFields)
}

func (s *CategoryService) DeleteCategory(ctx context.Context, categoryId string) error {
	convertedCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		return errs.BadRequest("INVALID_CATEGORY_ID", err)
	}
	return s.repository.DeleteCategory(ctx, convertedCategoryId)
}
