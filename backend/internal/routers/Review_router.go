package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewReviewRoutes(mainRouter *gin.RouterGroup, reviewHandler *handlers.ReviewHandler) {
	reviewRoute := mainRouter.Group("/review")

	reviewRoute.POST("/add", reviewHandler.AddNewReview)
	reviewRoute.PUT("/update/:id", reviewHandler.UpdateReview)
	reviewRoute.DELETE("/remove/:id", reviewHandler.RemoveReview)
	reviewRoute.GET("/:id", reviewHandler.GetReviewByID)

}
