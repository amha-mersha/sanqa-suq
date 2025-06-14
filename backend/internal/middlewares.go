package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*AppError); ok {
				c.JSON(appErr.StatusCode, gin.H{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
				return
			}

			// Unknown/internal error fallback
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}
