package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewAddressRoutes(router *gin.RouterGroup, addressHandler *handlers.AddressHandler, authMiddleware *middlewares.AuthMiddleware) {
	addresses := router.Group("/address")
	addresses.POST("", authMiddleware.AuthMiddleware(), addressHandler.CreateAddress)
	addresses.GET("", authMiddleware.AuthMiddleware(), addressHandler.GetUserAddresses)
	addresses.GET("/:id", authMiddleware.AuthMiddleware(), addressHandler.GetAddressByID)
	addresses.PUT("/:id", authMiddleware.AuthMiddleware(), addressHandler.UpdateAddress)
	addresses.DELETE("/:id", authMiddleware.AuthMiddleware(), addressHandler.DeleteAddress)
}
