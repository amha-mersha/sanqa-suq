package dtos

type AddressType string

const (
	ShippingAddress AddressType = "shipping"
	BillingAddress  AddressType = "billing"
)

type CreateAddressRequestDTO struct {
	Street     string      `json:"street" binding:"required"`
	City       string      `json:"city" binding:"required"`
	State      string      `json:"state"`
	PostalCode string      `json:"postal_code" binding:"required"`
	Country    string      `json:"country" binding:"required"`
	Type       AddressType `json:"type" binding:"required,oneof=shipping billing"`
}

type UpdateAddressRequestDTO struct {
	Street     string      `json:"street"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	PostalCode string      `json:"postal_code"`
	Country    string      `json:"country"`
	Type       AddressType `json:"type" binding:"omitempty,oneof=shipping billing"`
}

type AddressResponseDTO struct {
	AddressID  int         `json:"address_id"`
	Street     string      `json:"street"`
	City       string      `json:"city"`
	State      string      `json:"state"`
	PostalCode string      `json:"postal_code"`
	Country    string      `json:"country"`
	Type       AddressType `json:"type"`
}
