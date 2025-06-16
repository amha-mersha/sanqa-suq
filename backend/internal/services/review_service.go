package services

import (
	"context"
	"errors"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type ReviewService struct {
	repository *repositories.ReviewRepository
}

func (service *ReviewService) AddNewReview(ctx context.Context, reviewDTO *dtos.CreateReviewDTO) (*models.Review, error) {

	if reviewDTO.UserID == "" {
		return nil, errs.BadRequest("REVIEW_USER_ID_REQUIRED", nil)
	}
	if reviewDTO.ProductID == 0 {
		return nil, errs.BadRequest("REVIEW_PRODUCT_ID_REQUIRED", nil)
	}
	if reviewDTO.Rating < 1 || reviewDTO.Rating > 5 {
		return nil, errs.BadRequest("REVIEW_RATING_INVALID", nil)
	}

	return service.repository.CreateNewReview(ctx, reviewDTO)
}

func (service *ReviewService) UpdateReview(ctx context.Context, reviewID string, dto *dtos.UpdateReviewDTO) (*models.Review, error) {
	updateFields := make(map[string]any)

	if dto.Rating != nil {
		if *dto.Rating < 1 || *dto.Rating > 5 {
			return nil, errs.BadRequest("INVALID_REVIEW_RATING", nil)
		}
		updateFields["rating"] = *dto.Rating
	}

	if dto.Comment != nil {
		updateFields["comment"] = *dto.Comment
	}

	if len(updateFields) == 0 {
		return nil, errs.BadRequest("NO_FIELDS_TO_UPDATE", errors.New("no valid fields provided for update"))
	}

	return service.repository.UpdateReview(ctx, reviewID, dto)
}

func (service *ReviewService) GetReview(ctx context.Context, reviewID string) (*models.Review, error) {
	if reviewID == "" {
		return nil, errs.BadRequest("INVALID_REVIEW_ID", nil)
	}
	return service.repository.FindReviewByID(ctx, reviewID)
}

func (service *ReviewService) RemoveReview(ctx context.Context, reviewID string) error {
	if reviewID == "" {
		return errs.BadRequest("INVALID_REVIEW_ID", nil)
	}
	return service.repository.DeleteReviewByID(ctx, reviewID)
}
