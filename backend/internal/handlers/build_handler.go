package handlers

import (
	"net/http"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	"github.com/amha-mersha/sanqa-suq/internal/dtos"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/amha-mersha/sanqa-suq/internal/middlewares"
	"github.com/amha-mersha/sanqa-suq/internal/services"
	"github.com/gin-gonic/gin"
)

type BuildHandler struct {
	buildService *services.BuildService
}

func NewBuildHandler(buildService *services.BuildService) *BuildHandler {
	return &BuildHandler{
		buildService: buildService,
	}
}

func (h *BuildHandler) CreateBuild(c *gin.Context) {
	var req dtos.CreateBuildRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}

	// Get user ID from context
	claims, ok := c.Request.Context().Value(middlewares.UserClaimsKey).(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("MISSING_CLAIMS", nil))
		return
	}

	response, err := h.buildService.CreateBuild(c.Request.Context(), claims.UserID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *BuildHandler) GetUserBuilds(c *gin.Context) {
	claims, ok := c.Request.Context().Value(middlewares.UserClaimsKey).(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("MISSING_CLAIMS", nil))
		return
	}

	builds, err := h.buildService.GetUserBuilds(c.Request.Context(), claims.UserID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, builds)
}

func (h *BuildHandler) GetBuildByID(c *gin.Context) {
	buildID := c.Param("id")
	if buildID == "" {
		c.Error(errs.BadRequest("MISSING_BUILD_ID", nil))
		return
	}

	build, err := h.buildService.GetBuildByID(c.Request.Context(), buildID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, build)
}

func (h *BuildHandler) UpdateBuild(c *gin.Context) {
	buildID := c.Param("id")
	if buildID == "" {
		c.Error(errs.BadRequest("MISSING_BUILD_ID", nil))
		return
	}

	var req dtos.UpdateBuildRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}

	// Get user ID from context
	claims, ok := c.Request.Context().Value(middlewares.UserClaimsKey).(*auth.CustomClaims)
	if !ok {
		c.Error(errs.Unauthorized("MISSING_CLAIMS", nil))
		return
	}

	response, err := h.buildService.UpdateBuild(c.Request.Context(), buildID, claims.UserID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *BuildHandler) GetCompatibleProducts(c *gin.Context) {
	var req dtos.CompatibleProductsRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errs.BadRequest("INVALID_REQUEST_BODY", err))
		return
	}

	products, err := h.buildService.GetCompatibleProducts(c.Request.Context(), req.CategoryID, req.SelectedItems)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, products)
}
