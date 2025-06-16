package handlers

import (
	"net/http"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (handler *ProductHandler) AddNewProduct(ctx *gin.Context) {
	var productCreationDTO dtos.CreateProductDTO
	if err := ctx.ShouldBindBodyWithJSON(&productCreationDTO); err != nil {
		ctx.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}
	newProduct, err := handler.service.AddNewProduct(ctx, &productCreationDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"error": "PRODUCT_ADDED_SUCCESSFULLY", "data": struct {
		newProduct any `json:"new_product"`
	}{newProduct: newProduct}})
}

func (handler *ProductHandler) RemoveProduct(ctx *gin.Context) {
	param := ctx.Param("id")
	productId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_PRODUCT_ID"})
		return
	}
	err := handler.service.RemoveProduct(ctx, productId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "PRODUCT_REMOVED_SUCCESSFULLY"})
}

func (handler *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var productUpdateDTO dtos.ProductUpdateDTO
	if err := ctx.ShouldBindBodyWithJSON(&productUpdateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY"})
		return
	}
	productId, errConv := strconv.Atoi(ctx.Param("id"))
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_PRODUCT_ID"})
		return
	}
	err := handler.service.UpdateProduct(ctx, productId, &productUpdateDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "PRODUCT_UPDATED_SUCCESSFULLY"})
}

func (handler *ProductHandler) GetProduct(ctx *gin.Context) {
	param := ctx.Param("id")
	productId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.Error(errs.BadRequest("INVALID_PRODUCT_ID", errConv))
		return
	}
	product, err := handler.service.GetProduct(ctx, productId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "PRODUCT_FETCHED_SUCCESSFULLY", "data": struct {
		Product models.Products `json:"product"`
	}{Product: *product},
	})

}

func (handler *ProductHandler) GetAllProducts(ctx *gin.Context) {
	products, err := handler.service.GetAllProducts(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "PRODUCTS_FETCHED_SUCCESSFULLY", "data": struct {
		Products []*models.Products `json:"products"`
	}{Products: products}})
}

func (handler *ProductHandler) GetProductsByCategoryID(ctx *gin.Context) {
	category := ctx.Param("category_id")
	categoryId, errConv := strconv.Atoi(category)
	if errConv != nil {
		ctx.Error(errs.BadRequest("INVALID_CATEGORY_ID", errConv))
		return
	}
	products, err := handler.service.GetProductsByCategoryID(ctx, categoryId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "PRODUCTS_FETCHED_SUCCESSFULLY", "data": struct {
		Products []*models.Products `json:"products"`
	}{Products: products},
	})
}
func (handler *ProductHandler) AddNewReview(ctx *gin.Context) {
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

func (handler *ProductHandler) UpdateReview(ctx *gin.Context) {
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

func (handler *ProductHandler) RemoveReview(ctx *gin.Context) {
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

func (handler *ProductHandler) GetReviewByID(ctx *gin.Context) {
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
