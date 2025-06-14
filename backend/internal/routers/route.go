package routers

import (
	"fmt"

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
	apiRouter := rtr.Group(fmt.Sprintf("/api/%s/", config.Version))

	prodRepo := repositories.NewProductRepository(db)
	prodService := services.NewProductService(prodRepo)
	prodHandler := handlers.NewProductHandler(prodService)
	NewProductRoutes(apiRouter, prodHandler)

	catRepo := repositories.NewCategoryRepository(db)
	catService := services.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catService)
	NewCategoriesRoutes(apiRouter, catHandler)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	NewUserRoutes(apiRouter, userHandler)

	return nil
}
