package models

import "time"

type User struct {
	ID           string    `json:"user_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Phone        string    `json:"phone"`
	Role         string    `json:"role"`
	Provider     string    `json:"provider"`
	ProviderID   *string   `json:"provider_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

var UserRoles = []string{"customer", "seller", "admin"}
var UserProviders = []string{"local", "google"}
