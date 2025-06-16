package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/database"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	DB *database.DB
}

func NewOrderRepository(db *database.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// FetchAllOrders returns every order in the DB
func (r *OrderRepository) FetchAllOrders(ctx context.Context) ([]*models.Order, error) {
	query := `
        SELECT order_id, user_id, address_id, order_date, status, total_amount, payment_method, payment_date
        FROM orders
    `
	rows, err := r.DB.Pool.Query(ctx, query)
	if err != nil {
		return nil, errs.InternalError("failed to fetch orders", err)
	}
	defer rows.Close()

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByPos[*models.Order])
	if err != nil {
		return nil, errs.InternalError("failed to collect orders", err)
	}
	return orders, nil
}

// FindOrderByID returns one order by its UUID
func (r *OrderRepository) FindOrderByID(ctx context.Context, orderID string) (*models.Order, error) {
	var order models.Order
	query := `
        SELECT order_id, user_id, address_id, order_date, status, total_amount, payment_method, payment_date
        FROM orders
        WHERE order_id = $1
    `
	err := r.DB.Pool.QueryRow(ctx, query, orderID).Scan(
		&order.OrderID,
		&order.UserID,
		&order.AddressID,
		&order.OrderDate,
		&order.Status,
		&order.TotalAmount,
		&order.PaymentMethod,
		&order.PaymentDate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("order with id %s not found", orderID), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to find order with id %s", orderID), err)
	}
	return &order, nil
}

// FetchOrdersByUserID returns all orders placed by a given user
func (r *OrderRepository) FetchOrdersByUserID(ctx context.Context, userID string) ([]*models.Order, error) {
	query := `
        SELECT order_id, user_id, address_id, order_date, status, total_amount, payment_method, payment_date
        FROM orders
        WHERE user_id = $1
    `
	rows, err := r.DB.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to fetch orders for user %s", userID), err)
	}
	defer rows.Close()

	orders, err := pgx.CollectRows(rows, pgx.RowToStructByPos[*models.Order])
	if err != nil {
		return nil, errs.InternalError(fmt.Sprintf("failed to collect orders for user %s", userID), err)
	}
	if len(orders) == 0 {
		return nil, errs.NotFound(fmt.Sprintf("no orders found for user %s", userID), nil)
	}
	return orders, nil
}

// InsertNewOrder creates a new order record
func (r *OrderRepository) InsertNewOrder(ctx context.Context, dto *dtos.CreateOrderDTO) (*models.Order, error) {
	query := `
        INSERT INTO orders (
            user_id,
            address_id,
            total_amount,
            payment_method
        ) VALUES ($1, $2, $3, $4)
        RETURNING order_id, user_id, address_id, order_date, status, total_amount, payment_method, payment_date
    `
	newOrder := &models.Order{}
	err := r.DB.Pool.QueryRow(
		ctx, query,
		dto.UserID,
		dto.AddressID,
		dto.TotalAmount,
		dto.PaymentMethod,
	).Scan(
		&newOrder.OrderID,
		&newOrder.UserID,
		&newOrder.AddressID,
		&newOrder.OrderDate,
		&newOrder.Status,
		&newOrder.TotalAmount,
		&newOrder.PaymentMethod,
		&newOrder.PaymentDate,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23503" {
			return nil, errs.BadRequest("FOREIGN_KEY_VIOLATION", fmt.Errorf("invalid user_id or address_id: %w", err))
		}
		return nil, errs.InternalError("failed to insert new order", err)
	}
	return newOrder, nil
}

// UpdateOrder performs a partial update, returning the updated row
func (r *OrderRepository) UpdateOrder(ctx context.Context, orderID string, dto *dtos.UpdateOrderDTO) (*models.Order, error) {
	fields := make(map[string]any)
	if dto.Status != nil {
		fields["status"] = *dto.Status
	}
	if dto.TotalAmount != nil {
		fields["total_amount"] = *dto.TotalAmount
	}
	if dto.PaymentMethod != nil {
		fields["payment_method"] = *dto.PaymentMethod
	}
	if dto.PaymentDate != nil {
		fields["payment_date"] = *dto.PaymentDate
	}

	if len(fields) == 0 {
		return nil, errs.BadRequest("NO_FIELDS_TO_UPDATE", errors.New("no fields provided"))
	}

	setClauses := []string{}
	args := []any{}
	i := 1
	for col, val := range fields {
		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", col, i))
		args = append(args, val)
		i++
	}
	args = append(args, orderID)

	query := fmt.Sprintf(`
        UPDATE orders
        SET %s
        WHERE order_id = $%d
        RETURNING order_id, user_id, address_id, order_date, status, total_amount, payment_method, payment_date
    `, strings.Join(setClauses, ", "), i)

	updated := &models.Order{}
	err := r.DB.Pool.QueryRow(ctx, query, args...).Scan(
		&updated.OrderID,
		&updated.UserID,
		&updated.AddressID,
		&updated.OrderDate,
		&updated.Status,
		&updated.TotalAmount,
		&updated.PaymentMethod,
		&updated.PaymentDate,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFound(fmt.Sprintf("order with id %s not found", orderID), err)
		}
		return nil, errs.InternalError(fmt.Sprintf("failed to update order %s", orderID), err)
	}
	return updated, nil
}

// DeleteOrderByID removes an order
func (r *OrderRepository) DeleteOrderByID(ctx context.Context, orderID string) error {
	query := `DELETE FROM orders WHERE order_id = $1`
	result, err := r.DB.Pool.Exec(ctx, query, orderID)
	if err != nil {
		return errs.InternalError(fmt.Sprintf("failed to delete order %s", orderID), err)
	}
	if result.RowsAffected() == 0 {
		return errs.NotFound(fmt.Sprintf("order with id %s not found", orderID), nil)
	}
	return nil
}
