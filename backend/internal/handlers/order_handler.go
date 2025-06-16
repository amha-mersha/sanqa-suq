package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// AddNewOrder handles POST /order/add
func (handler *OrderHandler) AddNewOrder(ctx *gin.Context) {
	var dto dtos.CreateOrderDTO
	if err := ctx.ShouldBindBodyWithJSON(&dto); err != nil {
		ctx.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}
	order, err := handler.service.AddNewOrder(ctx, &dto)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "ORDER_CREATED_SUCCESSFULLY",
		"data": struct {
			Order models.Order `json:"order"`
		}{Order: *order},
	})
}

// GetOrder handles GET /order/:id
func (handler *OrderHandler) GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	if orderID == "" {
		ctx.Error(errs.BadRequest("INVALID_ORDER_ID", nil))
		return
	}
	order, err := handler.service.GetOrder(ctx, orderID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ORDER_FETCHED_SUCCESSFULLY",
		"data": struct {
			Order models.Order `json:"order"`
		}{Order: *order},
	})
}

// GetAllOrders handles GET /orders
func (handler *OrderHandler) GetAllOrders(ctx *gin.Context) {
	orders, err := handler.service.GetAllOrders(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ORDERS_FETCHED_SUCCESSFULLY",
		"data": struct {
			Orders []*models.Order `json:"orders"`
		}{Orders: orders},
	})
}

// GetOrdersByUser handles GET /order/user/:user_id
func (handler *OrderHandler) GetOrdersByUser(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.Error(errs.BadRequest("INVALID_USER_ID", nil))
		return
	}
	orders, err := handler.service.GetOrdersByUser(ctx, userID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ORDERS_FETCHED_BY_USER_SUCCESSFULLY",
		"data": struct {
			Orders []*models.Order `json:"orders"`
		}{Orders: orders},
	})
}

// UpdateOrder handles PUT /order/update/:id
func (handler *OrderHandler) UpdateOrder(ctx *gin.Context) {
	var dto dtos.UpdateOrderDTO
	if err := ctx.ShouldBindBodyWithJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY"})
		return
	}
	orderID := ctx.Param("id")
	if orderID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_ORDER_ID"})
		return
	}
	updated, err := handler.service.UpdateOrder(ctx, orderID, &dto)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ORDER_UPDATED_SUCCESSFULLY",
		"data": struct {
			Order models.Order `json:"order"`
		}{Order: *updated},
	})
}

// RemoveOrder handles DELETE /order/remove/:id
func (handler *OrderHandler) RemoveOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	if orderID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_ORDER_ID"})
		return
	}
	err := handler.service.RemoveOrder(ctx, orderID)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ORDER_REMOVED_SUCCESSFULLY"})
}
