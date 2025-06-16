package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) UserRegister(ctx *gin.Context) {
	var userRegisterDTO dtos.UserRegisterDTO
	if err := ctx.ShouldBindBodyWithJSON(&userRegisterDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY", "details": err.Error()})
		return
	}
	user, err := h.service.RegisterUser(ctx, &userRegisterDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "USER_REGISTERED_SUCCESSFULLY",
		"user_id": user.ID,
		"user":    user,
	})
}

func (h *UserHandler) UserLogin(ctx *gin.Context) {
	var userLoginDTO dtos.UserLoginDTO
	if err := ctx.ShouldBindBodyWithJSON(&userLoginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY", "details": err.Error()})
		return
	}
	token, user, err := h.service.LoginUser(ctx, &userLoginDTO)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.SetCookie("token", token, 3600*24, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	var userUpdateDTO dtos.UserUpdateDTO
	if err := ctx.ShouldBindBodyWithJSON(&userUpdateDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "INVALID_REQUEST_BODY", "details": err.Error()})
		return
	}

	// Get user ID from path parameter
	userId := ctx.Param("user_id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "MISSING_USER_ID", "details": "user ID is required"})
		return
	}

	// Get claims from context
	claims, exists := ctx.Get(string(middlewares.UserClaimsKey))
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "MISSING_CLAIMS", "details": "user claims not found"})
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "INVALID_CLAIMS", "details": "invalid user claims format"})
		return
	}

	// Check if user is updating their own profile
	if userClaims.UserID != userId {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "FORBIDDEN", "details": "you can only update your own profile"})
		return
	}

	updatedUser, err := h.service.UpdateUser(ctx, userId, &userUpdateDTO)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "USER_UPDATED_SUCCESSFULLY",
		"user":    updatedUser,
	})
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	userId := ctx.Param("user_id")
	user, err := h.service.GetUserById(ctx, userId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
