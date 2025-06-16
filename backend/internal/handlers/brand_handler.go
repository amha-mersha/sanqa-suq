package handlers

import (
	"net/http"
	"strconv"

	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/models"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	serivce *services.BrandService
}

func NewBrandHandler(service *services.BrandService) *BrandHandler {
	return &BrandHandler{
		serivce: service,
	}
}
func (h *BrandHandler) AddNewBrand(c *gin.Context) {
	var brandRequest dtos.CreateBrandRequest
	if err := c.ShouldBindJSON(&brandRequest); err != nil {
		c.Error(errs.BadRequest("Invalid request body", err))
		return
	}
	brand, err := h.serivce.CreateBrand(c.Request.Context(), &brandRequest)
	if err != nil {
		c.Error(errs.InternalError("Failed to create brand", err))
		return
	}
	c.JSON(201, gin.H{
		"data": struct {
			Brand models.Brands `json:"brand"`
		}{
			Brand: *brand,
		},
		"message": "Brand created successfully",
	})

}
func (h *BrandHandler) UpdateBrand(c *gin.Context) {
	// Logic to update an existing brand
}
func (h *BrandHandler) RemoveBrand(c *gin.Context) {
	// Logic to remove a brand
}
func (h *BrandHandler) GetBrand(c *gin.Context) {
	brandID := c.Param("id")
	convertedBrandID, err := strconv.Atoi(brandID)
	if err != nil {
		c.Error(errs.BadRequest("INVALID_BRAND_ID", err))
		return
	}

	brand, err := h.serivce.GetBrandByID(c.Request.Context(), convertedBrandID)
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, gin.H{
		"message": "BRAND_RETRIEVED_SUCCESSFULLY",
		"data": struct {
			Brand models.Brands `json:"brand"`
		}{
			Brand: *brand,
		},
	})
}

func (h *BrandHandler) GetAllBrands(c *gin.Context) {
	brands, err := h.serivce.GetBrands(c)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Brands retrieved successfully",
		"data": struct {
			Brands []*models.Brands `json:"brands"`
		}{
			Brands: brands,
		},
	})
}
