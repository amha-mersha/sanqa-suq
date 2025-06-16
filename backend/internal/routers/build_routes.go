package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewBuildRoutes(router *gin.RouterGroup, buildHandler *handlers.BuildHandler, authMiddleware *middlewares.AuthMiddleware) {
	builds := router.Group("/build")
	{
		// builds.POST("", authMiddleware.AuthMiddleware(), buildHandler.CreateBuild)
		// builds.GET("", authMiddleware.AuthMiddleware(), buildHandler.GetUserBuilds)
		// builds.GET("/:id", buildHandler.GetBuildByID)
		// builds.PUT("/:id", authMiddleware.AuthMiddleware(), buildHandler.UpdateBuild)

		builds.POST("", buildHandler.CreateBuild)
		builds.GET("", buildHandler.GetUserBuilds)
		builds.GET("/:id", buildHandler.GetBuildByID)
		builds.PUT("/:id", buildHandler.UpdateBuild)
	}
}
