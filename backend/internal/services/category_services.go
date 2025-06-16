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
	convCtgryId, err := strconv.Atoi(categoryId)
	if err != nil {
		return nil, errs.BadRequest("INVALID_CATEGORY_ID", err)
	}
	category, err := s.repository.GetCategoryById(ctx, convCtgryId)
	return category, err
}

func (s *CategoryService) GetCategoryWithChildren(ctx context.Context, categoryId string, limit int) ([]*models.CategoryNode, error) {
	convertedCategoryId, err := strconv.Atoi(categoryId)
	if err != nil {
		return nil, errs.BadRequest("INVALID_CATEGORY_ID", err)
	}
	return s.repository.FetchCategoryTree(ctx, convertedCategoryId, limit)
}

func (s *CategoryService) UpdateCategory(ctx context.Context, categoryId string, updateData *dtos.UpdateCategoryDTO) error {
	// check which fields are being updated and update only those
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
		// validate there is no loop in the requested update
		descendants, err := s.repository.FetchCategoryDescendants(ctx, convertedCategoryId)
		if err != nil {
			return err
		}
		for _, descendant := range descendants {
			if descendant.CategoryID == *updateData.ParentCategoryID {
				return errs.BadRequest("PARENT_CATEGORY_CANNOT_BE_A_DESCENDANT", nil)
			}
		}

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
