package services

import (
	"context"
	"errors"
	"slices"

	"github.com/amha-mersha/sanqa-suq/internal"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, userRegisterDTO *dtos.UserRegisterDTO) error {
	//Checking for valid User
	if !internal.ValidatePassword(userRegisterDTO.Password) {
		return internal.BadRequest("INVALID_PASSWORD", errors.New("password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character"))
	}
	//check email format and uniqueness
	if !internal.ValidateEmail(userRegisterDTO.Email) {
		return internal.BadRequest("INVALID_EMAIL", errors.New("email format is invalid"))
	}
	existingUser, errExisiting := s.repository.FindUserByEmail(ctx, userRegisterDTO.Email)
	if errExisiting == nil && existingUser != nil {
		return internal.Conflict("EMAIL_ALREADY_EXISTS", errors.New("email is already registered"))
	}
	if !internal.ValidatePhoneNumber(userRegisterDTO.Phone) {
		return internal.BadRequest("INVALID_PHONE", errors.New("phone number format is invalid"))
	}
	if slices.Contains(models.UserRoles, userRegisterDTO.Role) == false {
		return internal.BadRequest("INVALID_ROLE", errors.New("role must be one of the predefined roles"))
	}
	// Check if provider is valid and if provider ID is provided for non-local providers
	if slices.Contains(models.UserProviders, userRegisterDTO.Provider) == false {
		return internal.BadRequest("INVALID_PROVIDER", errors.New("provider must be one of the predefined providers"))
	}
	if userRegisterDTO.Provider != "local" && userRegisterDTO.ProviderID == "" {
		return internal.BadRequest("MISSING_PROVIDER_ID", errors.New("provider ID is required for non-local providers"))
	}

	hashedPassword, err := internal.HashPassword(userRegisterDTO.Password)
	if err != nil {
		return err
	}
	newUser := &models.User{
		FirstName:    userRegisterDTO.FirstName,
		LastName:     userRegisterDTO.LastName,
		Email:        userRegisterDTO.Email,
		PasswordHash: hashedPassword,
		Phone:        userRegisterDTO.Phone,
		Role:         userRegisterDTO.Role,
		Provider:     userRegisterDTO.Provider,
		ProviderID:   &userRegisterDTO.ProviderID,
	}
	errInsert := s.repository.InsertUser(ctx, newUser)
	return errInsert
}
