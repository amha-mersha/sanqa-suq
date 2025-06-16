package handlers

import "github.com/gin-gonic/gin"

func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  "UP",
		"message": "Service is running",
	})
}

func HealthPingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status":  "UP",
		"message": "Pong",
	})
}
