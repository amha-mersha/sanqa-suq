package main

import (
	"log"

	"github.com/amha-mersha/sanqa-suq/internal/configs"
	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	configs, errConfig := configs.LoadConfig(".env")
	if errConfig != nil {
		log.Fatal(errConfig)
	}
	// Initialize database
	db, err := database.NewDatabase(configs.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	r := gin.Default()
	errRoute := routers.NewRoute(configs, r)
	if errRoute != nil {
		log.Fatal(errRoute)
	}
	r.Run()
}
