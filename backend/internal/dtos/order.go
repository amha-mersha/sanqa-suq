package dtos

type CreateOrderDTO struct {
	UserID        string  `json:"user_id" binding:"required,uuid"`
	AddressID     int     `json:"address_id" binding:"required"`
	TotalAmount   float64 `json:"total_amount" binding:"required,gte=0"`
	PaymentMethod string  `json:"payment_method" binding:"required,oneof=cash credit_card debit_card"`
}

type UpdateOrderDTO struct {
	Status        *string  `json:"status" binding:"omitempty,oneof=pending shipped delivered cancelled"`
	TotalAmount   *float64 `json:"total_amount" binding:"omitempty,gte=0"`
	PaymentMethod *string  `json:"payment_method" binding:"omitempty,oneof=cash credit_card debit_card"`
	PaymentDate   *string  `json:"payment_date" binding:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}
