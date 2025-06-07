package dtos

type UserRegisterDTO struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Role       string `json:"role" binding:"required"`
	Provider   string `json:"provider" binding:"required"`
	ProviderId string `json:"provider_id"`
}
