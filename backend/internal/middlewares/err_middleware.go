package middlewares

import (
	"net/http"

	errs "github.com/amha-mersha/sanqa-suq/internal/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			if appErr, ok := err.(*errs.AppError); ok {
				c.JSON(appErr.StatusCode, gin.H{
					"error": appErr.Message,
					"code":  appErr.Code,
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}
