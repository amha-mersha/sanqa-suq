package handlers

import (
	"net/http"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
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
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "INVALID_PRODUCT_DATA", err, nil))
		return
	}
	err := handler.service.AddNewProduct(ctx, &productCreationDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "FAILED_TO_ADD_PRODUCT", err, nil))
		return
	}
	ctx.JSON(http.StatusCreated, NewResponseJsonStruct(StatusSuccess, "PRODUCT_ADDED_SUCCESSFULLY", nil, nil))
}

func (handler *ProductHandler) RemoveProduct(ctx *gin.Context) {
	param := ctx.Param("id")
	productId, errConv := strconv.Atoi(param)
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "INVALID_PRODUCT_ID", errConv, nil))
		return
	}
	err := handler.service.RemoveProduct(ctx, productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "FAILED_TO_REMOVE_PRODUCT", err, nil))
		return
	}
	ctx.JSON(http.StatusOK, NewResponseJsonStruct(StatusSuccess, "PRODUCT_REMOVED_SUCCESSFULLY", nil, nil))
}

func (handler *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var productUpdateDTO dtos.ProductUpdateDTO
	if err := ctx.ShouldBindBodyWithJSON(&productUpdateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "INVALID_PRODUCT_DATA", err, nil))
		return
	}
	productId, errConv := strconv.Atoi(ctx.Param("id"))
	if errConv != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "INVALID_PRODUCT_ID", errConv, nil))
		return
	}
	err := handler.service.UpdateProduct(ctx, productId, &productUpdateDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "FAILED_TO_UPDATE_PRODUCT", err, nil))
		return
	}
	ctx.JSON(http.StatusOK, NewResponseJsonStruct(StatusSuccess, "PRODUCT_UPDATED_SUCCESSFULLY", nil, nil))
}
