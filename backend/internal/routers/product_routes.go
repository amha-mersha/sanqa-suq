package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewProductRoutes(mainRouter *gin.Engine, productHanlder *handlers.ProductHandler) {
	productRoute := mainRouter.Group("/products")

	productRoute.POST("/add", productHanlder.AddNewProduct)
	productRoute.PUT("/update:id", productHanlder.UpdateProduct)
	productRoute.DELETE("/remove:id", productHanlder.RemoveProduct)
}
