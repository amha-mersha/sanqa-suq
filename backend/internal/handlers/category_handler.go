package handlers

import (
	"net/http"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	serivce *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		serivce: service,
	}
}

func (handler *CategoryHandler) GetAllCategroies(ctx *gin.Context) {
	categoryList, err := handler.serivce.GetCategoryList(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": categoryList})
}

func (handler *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var newCategory dtos.CreateCategoryDTO
	if errUnmarshal := ctx.ShouldBindBodyWithJSON(&newCategory); errUnmarshal != nil {
		ctx.Error(errs.BadRequest("INVALID_REQUEST_BODY", errUnmarshal))
		return
	}
	createdCategory, err := handler.serivce.CreateCategory(ctx, &newCategory)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "CATEGORY_CREATED_SUCCESSFULLY", "data": createdCategory})
}
func (h *CategoryHandler) GetCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.Error(errs.BadRequest("INVALID_CATEGORY_ID", nil))
		return
	}

	limitStr := ctx.DefaultQuery("limit", "-1")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < -1 {
		ctx.Error(errs.BadRequest("INVALID_LIMIT_VALUE", err))
		return
	}
	if limit == 0 {

		category, err := h.serivce.GetCategoryById(ctx, categoryId)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": category})
	} else {
		categories, err := h.serivce.GetCategoryWithChildren(ctx, categoryId, limit)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": categories})
	}
}
func (handler *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var updateData dtos.UpdateCategoryDTO
	if err := ctx.ShouldBindBodyWithJSON(&updateData); err != nil {
		ctx.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.Error(errs.BadRequest("INVALID_CATEGORY_ID", nil))
		return
	}
	err := handler.serivce.UpdateCategory(ctx, categoryId, &updateData)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "CATEGORY_UPDATED_SUCCESSFULLY"})
}
func (handler *CategoryHandler) DeleteCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.Error(errs.BadRequest("INVALID_CATEGORY_ID", nil))
		return
	}
	err := handler.serivce.DeleteCategory(ctx, categoryId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "CATEGORY_DELETED_SUCCESSFULLY"})
}
