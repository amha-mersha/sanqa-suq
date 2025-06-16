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

func (handler *ProductHandler) GetProductSpecifications(ctx *gin.Context) {
	param := ctx.Param("id")
	productId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.Error(errs.BadRequest("INVALID_PRODUCT_ID", errConv))
		return
	}
	product, err := handler.service.GetProductSpecifications(ctx, productId)
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
