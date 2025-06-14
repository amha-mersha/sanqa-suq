package dtos

type UserRegisterDTO struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	Phone      string `json:"phone" binding:"required"`
	Address    string `json:"address" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=customer seller admin"`
	Provider   string `json:"provider" binding:"required,oneof=local google"`
	ProviderID string `json:"provider_id" binding:"omitempty"`
}
