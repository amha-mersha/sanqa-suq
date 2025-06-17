package middlewares

import (
	"context"
	"slices"
	"strings"

	"github.com/amha-mersha/sanqa-suq/internal/auth"
	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService *auth.JWTService
}

func NewAuthMiddleware(jwtService *auth.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

type contextKey string

const UserClaimsKey contextKey = "user_claims"

// AuthMiddleware verifies JWT tokens in requests
func (a *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ""
		// Check cookie first
		if cookie, err := c.Cookie("token"); err == nil {
			tokenString = cookie
		} else {
			// Fallback to Authorization header
			authHeader := c.GetHeader("Authorization")
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			}
		}

		if tokenString == "" {
			c.Error(errs.Unauthorized("missing token", nil))
			c.Abort()
			return
		}

		claims, err := a.jwtService.ValidateToken(tokenString)
		if err != nil {
			c.Error(errs.Unauthorized("invalid token", err))
			c.Abort()
			return
		}

		// Set claims only in request context
		ctx := context.WithValue(c.Request.Context(), UserClaimsKey, claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// RequireRole creates a middleware that checks if the user has the required role
func (a *AuthMiddleware) RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get claims from context
		claims, exists := c.Request.Context().Value(UserClaimsKey).(*auth.CustomClaims)
		if !exists {
			c.Error(errs.Unauthorized("missing claims", nil))
			c.Abort()
			return
		}

		// Check if user has the required role
		if claims.Role != requiredRole {
			c.Error(errs.Forbidden("insufficient permissions", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAnyRole creates a middleware that checks if the user has any of the required roles
func (a *AuthMiddleware) RequireAnyRole(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get claims from context
		claims, exists := c.Request.Context().Value(UserClaimsKey).(*auth.CustomClaims)
		if !exists {
			c.Error(errs.Unauthorized("missing claims", nil))
			c.Abort()
			return
		}

		// Check if user has any of the required roles
		hasRole := slices.Contains(requiredRoles, claims.Role)

		if !hasRole {
			c.Error(errs.Forbidden("insufficient permissions", nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
