package routers

import (
	"fmt"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	"github.com/amha-mersha/sanqa-suq/internal/configs"
	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
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
	apiRouter.Use(middlewares.ErrorHandler())
	apiRouter.GET("health", handlers.HealthCheckHandler)
	apiRouter.GET("ping", handlers.HealthPingHandler)

	prodRepo := repositories.NewProductRepository(db)
	prodService := services.NewProductService(prodRepo)
	prodHandler := handlers.NewProductHandler(prodService)
	NewProductRoutes(apiRouter, prodHandler)

	catRepo := repositories.NewCategoryRepository(db)
	catService := services.NewCategoryService(catRepo)
	catHandler := handlers.NewCategoryHandler(catService)
	NewCategoriesRoutes(apiRouter, catHandler)

	userRepo := repositories.NewUserRepository(db)
	authService := auth.NewJWTService(config.JWTSecret, config.JWTIssuer)
	authMiddleware := middlewares.NewAuthMiddleware(authService)
	userService := services.NewUserService(userRepo, authService)
	userHandler := handlers.NewUserHandler(userService)
	NewUserRoutes(apiRouter, userHandler, authMiddleware)

	brandRepo := repositories.NewBrandRepository(db)
	brandService := services.NewBrandService(brandRepo)
	brandHandler := handlers.NewBrandHandler(brandService)
	NewBrandRoutes(apiRouter, brandHandler)

	return nil
}
