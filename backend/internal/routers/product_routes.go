package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewProductRoutes(mainRouter *gin.RouterGroup, productHanlder *handlers.ProductHandler) {
	productRoute := mainRouter.Group("/product")

	productRoute.POST("/add", productHanlder.AddNewProduct)
	productRoute.PUT("/update:id", productHanlder.UpdateProduct)
	productRoute.DELETE("/remove:id", productHanlder.RemoveProduct)
}
