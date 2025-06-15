package dtos

type UserRegisterDTO struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	Phone      string `json:"phone" binding:"required"`
	Role       string `json:"role" binding:"required,oneof=customer seller admin"`
	Provider   string `json:"provider" binding:"required,oneof=local google"`
	ProviderID string `json:"provider_id" binding:"omitempty"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserUpdateDTO struct {
	FirstName *string `json:"first_name" binding:"omitempty"`
	LastName  *string `json:"last_name" binding:"omitempty"`
	Phone     *string `json:"phone" binding:"omitempty"`
	Role      *string `json:"role" binding:"omitempty,oneof=customer seller admin"`
}
