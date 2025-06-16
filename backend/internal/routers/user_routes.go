package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(mainRouter *gin.RouterGroup, userHandler *handlers.UserHandler, middleware *middlewares.AuthMiddleware) {
	userRoute := mainRouter.Group("/user")

	userRoute.POST("/signup", userHandler.UserRegister)
	userRoute.POST("/login", userHandler.UserLogin)

	userRoute.Use(middleware.AuthMiddleware())

	userRoute.GET("/:user_id", userHandler.GetUserById)
	userRoute.PUT("/:user_id", userHandler.UpdateUser)
}
