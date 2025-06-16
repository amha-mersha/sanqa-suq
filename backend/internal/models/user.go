package models

import "time"

type User struct {
	ID           string
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	Phone        string
	Role         string
	Provider     string
	ProviderID   *string
	CreatedAt    time.Time
}

var UserRoles = []string{"customer", "seller", "admin"}
var UserProviders = []string{"local", "google"}
