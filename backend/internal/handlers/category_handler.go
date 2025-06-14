package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
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
	if errUnmarshal := ctx.ShouldBindBodyWithJSON(newCategory); errUnmarshal != nil {
		ctx.Error(internal.BadRequest("INVALID_REQUEST_BODY", errUnmarshal))
	}
	err := handler.serivce.CreateCategory(ctx, &newCategory)
	if err != nil {
		ctx.Error(err)
		return
	}
}
func (h *CategoryHandler) GetCategory(ctx *gin.Context) {
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.Error(internal.BadRequest("INVALID_CATEGORY_ID", nil))
		return
	}
	category, err := h.serivce.GetCategoryById(ctx, categoryId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": category})
}
func (handler *CategoryHandler) UpdateCategory(ctx *gin.Context) {
	var updateData dtos.UpdateCategoryDTO
	if err := ctx.ShouldBindBodyWithJSON(&updateData); err != nil {
		ctx.Error(internal.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}
	categoryId := ctx.Param("id")
	if categoryId == "" {
		ctx.Error(internal.BadRequest("INVALID_CATEGORY_ID", nil))
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
		ctx.Error(internal.BadRequest("INVALID_CATEGORY_ID", nil))
		return
	}
	err := handler.serivce.DeleteCategory(ctx, categoryId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "CATEGORY_DELETED_SUCCESSFULLY"})
}
