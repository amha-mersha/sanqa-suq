package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}
func (handler *ReviewHandler) AddNewReview(ctx *gin.Context) {
	var reviewDTO dtos.CreateReviewDTO

	if err := ctx.ShouldBindBodyWithJSON(&reviewDTO); err != nil {
		ctx.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}

	newReview, err := handler.service.AddNewReview(ctx, &reviewDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"error": "REVIEW_ADDED_SUCCESSFULLY",
		"data": struct {
			NewReview any `json:"new_review"`
		}{
			NewReview: newReview,
		},
	})
}

func (handler *ReviewHandler) UpdateReview(ctx *gin.Context) {
	var reviewUpdateDTO dtos.UpdateReviewDTO

	if err := ctx.ShouldBindBodyWithJSON(&reviewUpdateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY"})
		return
	}

	reviewID := ctx.Param("id")
	if reviewID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REVIEW_ID"})
		return
	}

	updatedReview, err := handler.service.UpdateReview(ctx, reviewID, &reviewUpdateDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "REVIEW_UPDATED_SUCCESSFULLY",
		"data":    updatedReview,
	})
}

func (handler *ReviewHandler) RemoveReview(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	if reviewID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REVIEW_ID"})
		return
	}

	err := handler.service.RemoveReview(ctx, reviewID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "REVIEW_REMOVED_SUCCESSFULLY"})
}

func (handler *ReviewHandler) GetReviewByID(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	if reviewID == "" {
		ctx.Error(errs.BadRequest("INVALID_REVIEW_ID", nil))
		return
	}

	review, err := handler.service.GetReview(ctx, reviewID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "REVIEW_FETCHED_SUCCESSFULLY",
		"data": struct {
			Review models.Review `json:"review"`
		}{
			Review: *review,
		},
	})
}
