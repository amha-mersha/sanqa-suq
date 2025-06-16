package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewBrandRoutes(mainRouter *gin.RouterGroup, brandHandler *handlers.BrandHandler) {
	brandRoute := mainRouter.Group("/brand")

	brandRoute.POST("/add", brandHandler.AddNewBrand)
	// brandRoute.PUT("/update/:id", brandHandler.UpdateBrand)
	// brandRoute.DELETE("/remove:id", brandHandler.RemoveBrand)
	brandRoute.GET("/:id", brandHandler.GetBrand)
	brandRoute.GET("", brandHandler.GetAllBrands)
}
