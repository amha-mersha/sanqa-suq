package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(mainRouter *gin.RouterGroup, userHandler *handlers.UserHandler, middleware *middlewares.AuthMiddleware) {
	userRoute := mainRouter.Group("/user")

	/*
		POST /api/users/register
		Data: { email, password, first_name, last_name, phone }
		Response: { user_id, email, role }
	*/
	userRoute.POST("/signup", userHandler.UserRegister)

	userRoute.Use(middleware.AuthMiddleware())
	/*
		POST /api/users/login
		Data: { email, password }
		Response: { token }
	*/
	userRoute.POST("/login", userHandler.UserLogin)

	/*
		GET /api/users/:user_id
		Response: { user_id, email, first_name, last_name, phone, role }
	*/
	userRoute.GET("users/:user_id", userHandler.GetUserById)

	/*
		PUT /api/users/:user_id
		Data: { first_name, last_name, phone }
		Response: { user_id, updated_fields }
	*/
	userRoute.POST("users/:user_id", userHandler.UpdateUser)
}
