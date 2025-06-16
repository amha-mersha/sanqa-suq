package models

import "time"

type Order struct {
	OrderID       string     `json:"order_id"`
	UserID        string     `json:"user_id"`
	AddressID     int        `json:"address_id"`
	OrderDate     time.Time  `json:"order_date"`
	Status        string     `json:"status"`
	TotalAmount   float64    `json:"total_amount"`
	PaymentMethod string     `json:"payment_method"`
	PaymentDate   *time.Time `json:"payment_date"` // optional
}
