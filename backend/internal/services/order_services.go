package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type OrderService struct {
	repository *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{repository: repo}
}

// Create a new order
func (service *OrderService) AddNewOrder(ctx context.Context, dto *dtos.CreateOrderDTO) (*models.Order, error) {
	if dto.UserID == "" {
		return nil, errs.BadRequest("ORDER_USER_ID_REQUIRED", nil)
	}
	if dto.AddressID <= 0 {
		return nil, errs.BadRequest("ORDER_ADDRESS_ID_REQUIRED", nil)
	}
	if dto.TotalAmount < 0 {
		return nil, errs.BadRequest("ORDER_TOTAL_AMOUNT_INVALID", nil)
	}
	// status and payment_date are set by DB defaults
	// payment method validity is handled by binding tag

	return service.repository.InsertNewOrder(ctx, dto)
}

// Update an existing order (partial updates)
func (service *OrderService) UpdateOrder(ctx context.Context, orderID string, dto *dtos.UpdateOrderDTO) (*models.Order, error) {
	updateFields := make(map[string]any)

	if dto.Status != nil {
		updateFields["status"] = *dto.Status
	}
	if dto.TotalAmount != nil {
		if *dto.TotalAmount < 0 {
			return nil, errs.BadRequest("ORDER_TOTAL_AMOUNT_INVALID", nil)
		}
		updateFields["total_amount"] = *dto.TotalAmount
	}
	if dto.PaymentMethod != nil {
		updateFields["payment_method"] = *dto.PaymentMethod
	}
	if dto.PaymentDate != nil {
		// parse or trust the incoming time value
		t, err := time.Parse(time.RFC3339, *dto.PaymentDate)
		if err != nil {
			return nil, errs.BadRequest("INVALID_PAYMENT_DATE_FORMAT", err)
		}
		updateFields["payment_date"] = t
	}

	if len(updateFields) == 0 {
		return nil, errs.BadRequest("NO_FIELDS_TO_UPDATE", errors.New("no valid fields provided for update"))
	}

	return service.repository.UpdateOrder(ctx, orderID, dto)
}

// Fetch one order by its ID
func (service *OrderService) GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	if orderID == "" {
		return nil, errs.BadRequest("INVALID_ORDER_ID", nil)
	}
	return service.repository.FindOrderByID(ctx, orderID)
}

// Fetch all orders
func (service *OrderService) GetAllOrders(ctx context.Context) ([]*models.Order, error) {
	orders, err := service.repository.FetchAllOrders(ctx)
	if err != nil {
		return nil, errs.InternalError("failed to fetch all orders", err)
	}
	return orders, nil
}

// Fetch orders for a specific user
func (service *OrderService) GetOrdersByUser(ctx context.Context, userID string) ([]*models.Order, error) {
	if userID == "" {
		return nil, errs.BadRequest("INVALID_USER_ID", nil)
	}
	orders, err := service.repository.FetchOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to fetch orders for user %s", userID), err)
	}
	return orders, nil
}

// Delete an order by its ID
func (service *OrderService) RemoveOrder(ctx context.Context, orderID string) error {
	if orderID == "" {
		return errs.BadRequest("INVALID_ORDER_ID", nil)
	}
	return service.repository.DeleteOrderByID(ctx, orderID)
}
