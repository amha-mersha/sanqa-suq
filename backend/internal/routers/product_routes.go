package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewProductRoutes(mainRouter *gin.Engine, productHanlder *handlers.ProductHandler) {
	productRoute := mainRouter.Group("/products")

	productRoute.POST("/add", ProductHandler.)
}
