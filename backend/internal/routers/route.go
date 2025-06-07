package routers

import "github.com/gin-gonic/gin"

func NewRoute() {
	r := gin.Default()
	NewUserRoutes(r)
	NewProductRoutes()
	NewOrderRoutes()
	NewAddressRoutes()

}
