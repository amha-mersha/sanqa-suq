package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewCategoriesRoutes(mainRouter *gin.RouterGroup, categoryHandler *handlers.CategoryHandler) {
	categoryRoutes := mainRouter.Group("/categories")

	categoryRoutes.GET("", categoryHandler.GetAllCategroies)
	categoryRoutes.POST("", categoryHandler.CreateCategory)
	categoryRoutes.GET(":id", categoryHandler.GetCategory)
	categoryRoutes.PUT(":id", categoryHandler.UpdateCategory)
	categoryRoutes.DELETE(":id", categoryHandler.DeleteCategory)
}
