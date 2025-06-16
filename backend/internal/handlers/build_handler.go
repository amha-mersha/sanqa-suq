package handlers

import (
	"net/http"

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
		c.Error(err)
		return
	}

	// Get user ID from context
	claims, exists := c.Request.Context().Value(middlewares.UserClaimsKey).(map[string]interface{})
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	response, err := h.buildService.CreateBuild(c.Request.Context(), userID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *BuildHandler) GetUserBuilds(c *gin.Context) {
	claims, exists := c.Request.Context().Value(middlewares.UserClaimsKey).(map[string]interface{})
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	builds, err := h.buildService.GetUserBuilds(c.Request.Context(), userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, builds)
}

func (h *BuildHandler) GetBuildByID(c *gin.Context) {
	buildID := c.Param("id")
	if buildID == "" {
		c.Error(errs.BadRequest("build ID is required", nil))
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
		c.Error(errs.BadRequest("build ID is required", nil))
		return
	}

	var req dtos.UpdateBuildRequestDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	// Get user ID from context
	claims, exists := c.Request.Context().Value(middlewares.UserClaimsKey).(map[string]interface{})
	if !exists {
		c.Error(errs.Unauthorized("user not authenticated", nil))
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		c.Error(errs.Unauthorized("invalid user claims", nil))
		return
	}

	response, err := h.buildService.UpdateBuild(c.Request.Context(), buildID, userID, &req)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
