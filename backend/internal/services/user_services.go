package services

import (
	"context"
	"errors"
	"slices"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/repositories"
	"github.com/amha-mersha/sanqa-suq/internal/utils"
)

type UserService struct {
	repository  *repositories.UserRepository
	authService *auth.JWTService
}

func NewUserService(repository *repositories.UserRepository, jwtService *auth.JWTService) *UserService {
	return &UserService{
		repository:  repository,
		authService: jwtService,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, userRegisterDTO *dtos.UserRegisterDTO) (*models.User, error) {
	//Checking for valid User
	if !utils.ValidatePassword(userRegisterDTO.Password) {
		return nil, errs.BadRequest("INVALID_PASSWORD", errors.New("password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character"))
	}
	//check email format and uniqueness
	if !utils.ValidateEmail(userRegisterDTO.Email) {
		return nil, errs.BadRequest("INVALID_EMAIL", errors.New("email format is invalid"))
	}
	existingUser, errExisiting := s.repository.FindUserByEmail(ctx, userRegisterDTO.Email)
	if errExisiting == nil && existingUser != nil {
		return nil, errs.Conflict("EMAIL_ALREADY_EXISTS", errors.New("email is already registered"))
	}
	if !utils.ValidatePhoneNumber(userRegisterDTO.Phone) {
		return nil, errs.BadRequest("INVALID_PHONE", errors.New("phone number format is invalid"))
	}
	if !slices.Contains(models.UserRoles, userRegisterDTO.Role) {
		return nil, errs.BadRequest("INVALID_ROLE", errors.New("role must be one of the predefined roles"))
	}
	// Check if provider is valid and if provider ID is provided for non-local providers
	if !slices.Contains(models.UserProviders, userRegisterDTO.Provider) {
		return nil, errs.BadRequest("INVALID_PROVIDER", errors.New("provider must be one of the predefined providers"))
	}
	if userRegisterDTO.Provider != "local" && userRegisterDTO.ProviderID == "" {
		return nil, errs.BadRequest("MISSING_PROVIDER_ID", errors.New("provider ID is required for non-local providers"))
	}

	hashedPassword, err := utils.HashPassword(userRegisterDTO.Password)
	if err != nil {
		return nil, err
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
	insertedUser, errInsert := s.repository.InsertUser(ctx, newUser)
	if errInsert != nil {
		return nil, errInsert
	}
	return insertedUser, nil
}

func (s *UserService) LoginUser(ctx context.Context, userLoginDTO *dtos.UserLoginDTO) (string, *models.User, error) {
	checkoutUser, err := s.repository.FindUserByEmail(ctx, userLoginDTO.Email)
	if err != nil || checkoutUser == nil {
		return "", nil, err
	}
	if !utils.ComparePasswords(userLoginDTO.Password, checkoutUser.PasswordHash) {
		return "", nil, errs.Unauthorized("INVALID_CREDENTIALS", errors.New("email or password is incorrect"))
	}
	token, err := s.authService.GenerateToken(checkoutUser.ID, checkoutUser.Role, checkoutUser.Email, checkoutUser.Provider, *checkoutUser.ProviderID)
	if err != nil {
		return "", nil, errs.InternalError("TOKEN_GENERATION_FAILED", err)
	}
	return token, checkoutUser, nil
}

func (s *UserService) UpdateUser(ctx context.Context, userId string, userUpdateDTO *dtos.UserUpdateDTO) (*models.User, error) {
	//start building the updateable fields map
	updateableFields := make(map[string]any)
	if userUpdateDTO.FirstName != nil {
		if *userUpdateDTO.FirstName == "" {
			return nil, errs.BadRequest("INVALID_FIRST_NAME", errors.New("first name cannot be empty"))
		}
		updateableFields["first_name"] = *userUpdateDTO.FirstName
	}
	if userUpdateDTO.LastName != nil {
		if *userUpdateDTO.LastName == "" {
			return nil, errs.BadRequest("INVALID_LAST_NAME", errors.New("last name cannot be empty"))
		}
		updateableFields["last_name"] = *userUpdateDTO.LastName
	}
	if userUpdateDTO.Phone != nil {
		if !utils.ValidatePhoneNumber(*userUpdateDTO.Phone) {
			return nil, errs.BadRequest("INVALID_PHONE", errors.New("phone number format is invalid"))
		}
		updateableFields["phone"] = *userUpdateDTO.Phone
	}
	if userUpdateDTO.Role != nil {
		if slices.Contains(models.UserRoles, *userUpdateDTO.Role) {
			updateableFields["role"] = userUpdateDTO.Role
		} else {
			return nil, errs.BadRequest("INVALID_ROLE", errors.New("role must be one of the predefined roles"))
		}
	}
	if len(updateableFields) == 0 {
		return nil, errs.BadRequest("NO_FIELDS_TO_UPDATE", errors.New("no valid fields provided for update"))
	}
	return s.repository.UpdateUser(ctx, userId, updateableFields)
}

func (s *UserService) GetUserById(ctx context.Context, userId string) (*models.User, error) {
	return s.repository.FindUserByID(ctx, userId)
}
