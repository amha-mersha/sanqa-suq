package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewUserRoutes(mainRouter *gin.Engine, userHandler *handlers.UserHandler) {
	userRoute := mainRouter.Group("/user")

	/*
		POST /api/users/register
		Data: { email, password, first_name, last_name, phone }
		Response: { user_id, email, role }
	*/
	userRoute.POST("/register", userHandler.UserRegister)

	/*
		POST /api/users/login
		Data: { email, password }
		Response: { token }
	*/
	userRoute.POST("/login", UserLogin)

	/*
		POST /api/users/oauth/google
		Data: { google_token }
		Response: { user_id, email, token }
	*/
	userRoute.POST("oauth/google", UserLogin)

	/*
		GET /api/users/:user_id
		Response: { user_id, email, first_name, last_name, phone, role }
	*/
	userRoute.GET("users/:user_id", GetUserById)

	/*
		PUT /api/users/:user_id
		Data: { first_name, last_name, phone }
		Response: { user_id, updated_fields }
	*/
	userRoute.POST("users/:user_id", UpdateUser)
}
