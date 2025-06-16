package routers

import (
	"github.com/amha-mersha/sanqa-suq/internal/handlers"
	"github.com/gin-gonic/gin"
)

// Register order-related routes onto the main router group
func NewOrderRoutes(mainRouter *gin.RouterGroup, orderHandler *handlers.OrderHandler) {
	// Group routes under /order
	order := mainRouter.Group("/order")
	{
		order.POST("/add", orderHandler.AddNewOrder)              // POST   /order/add
		order.GET(":id", orderHandler.GetOrder)                   // GET    /order/:id
		order.PUT("/update/:id", orderHandler.UpdateOrder)        // PUT    /order/update/:id
		order.DELETE("/remove/:id", orderHandler.RemoveOrder)     // DELETE /order/remove/:id
		order.GET("/user/:user_id", orderHandler.GetOrdersByUser) // GET    /order/user/:user_id
	}

	// Fetch all orders
	mainRouter.GET("/orders", orderHandler.GetAllOrders) // GET /orders
}
