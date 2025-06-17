package models

import (
	"time"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
)

type Address struct {
	AddressID  int              `json:"address_id" db:"address_id"`
	UserID     string           `json:"user_id" db:"user_id"`
	Street     string           `json:"street" db:"street"`
	City       string           `json:"city" db:"city"`
	State      string           `json:"state" db:"state"`
	PostalCode string           `json:"postal_code" db:"postal_code"`
	Country    string           `json:"country" db:"country"`
	Type       dtos.AddressType `json:"type" db:"address_type"`
	CreatedAt  time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at" db:"updated_at"`
}
