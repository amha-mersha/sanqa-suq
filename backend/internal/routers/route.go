package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/configs"
	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

func NewRoute(config *configs.Config, rtr *gin.Engine) error {
	db, err := database.NewDatabase(config.DatabaseUrl)
	if err != nil {
		return err
	}
	prodRepo := repositories.NewProductRepository(db)
	prodService := services.NewProductService(prodRepo)
	prodHandler := handlers.NewProductHandler(prodService)
	NewProductRoutes(rtr, prodHandler)
	return nil
}
