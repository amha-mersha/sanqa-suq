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
	productRoute.GET("/:id", productHanlder.GetProduct)
	productRoute.GET("", productHanlder.GetAllProducts)
	productRoute.GET("/get-by-category/:category_id", productHanlder.GetProductsByCategoryID)
}
