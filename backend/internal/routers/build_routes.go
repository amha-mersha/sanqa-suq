package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func NewBuildRoutes(router *gin.RouterGroup, buildHandler *handlers.BuildHandler, authMiddleware *middlewares.AuthMiddleware) {
	builds := router.Group("/build")
	builds.POST("", authMiddleware.AuthMiddleware(), buildHandler.CreateBuild)
	builds.GET("", authMiddleware.AuthMiddleware(), buildHandler.GetUserBuilds)
	builds.GET("/:id", authMiddleware.AuthMiddleware(), buildHandler.GetBuildByID)
	builds.PUT("/:id", authMiddleware.AuthMiddleware(), buildHandler.UpdateBuild)
	builds.POST("/compatible", authMiddleware.AuthMiddleware(), buildHandler.GetCompatibleProducts)
}
