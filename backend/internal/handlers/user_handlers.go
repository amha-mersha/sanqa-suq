package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
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
		ctx.JSON(http.StatusBadRequest, NewResponseJsonStruct(StatusError, "INVALID_USER_DATA", err, nil))
	}

	err := h.service.RegisterUser(ctx, userRegisterDTO)
}
func (h *UserHandler) UserLogin(ctx *gin.Context)   {}
func (h *UserHandler) UpdateUser(ctx *gin.Context)  {}
func (h *UserHandler) GetUserById(ctx *gin.Context) {}
