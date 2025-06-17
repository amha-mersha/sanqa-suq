package handlers

import (
	"net/http"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	addressService *services.AddressService
}

func NewAddressHandler(addressService *services.AddressService) *AddressHandler {
	return &AddressHandler{
		addressService: addressService,
	}
}

func (h *AddressHandler) CreateAddress(c *gin.Context) {
	var req dtos.CreateAddressRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest("invalid request body", err))
		return
	}

	claims, exists := c.Get(string(middlewares.UserClaimsKey))
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	address, err := h.addressService.CreateAddress(c.Request.Context(), userClaims.UserID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, address)
}

func (h *AddressHandler) GetAddressByID(c *gin.Context) {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errs.BadRequest("invalid address ID", err))
		return
	}

	claims, exists := c.Get(string(middlewares.UserClaimsKey))
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	address, err := h.addressService.GetAddressByID(c.Request.Context(), addressID, userClaims.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) GetUserAddresses(c *gin.Context) {
	claims, exists := c.Get(string(middlewares.UserClaimsKey))
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	addresses, err := h.addressService.GetUserAddresses(c.Request.Context(), userClaims.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errs.BadRequest("invalid address ID", err))
		return
	}

	var req dtos.UpdateAddressRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest("invalid request body", err))
		return
	}

	claims, exists := c.Get(string(middlewares.UserClaimsKey))
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	address, err := h.addressService.UpdateAddress(c.Request.Context(), addressID, userClaims.UserID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, address)
}

func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	addressID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errs.BadRequest("invalid address ID", err))
		return
	}

	claims, exists := c.Get(string(middlewares.UserClaimsKey))
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userClaims, ok := claims.(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	if err := h.addressService.DeleteAddress(c.Request.Context(), addressID, userClaims.UserID); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
