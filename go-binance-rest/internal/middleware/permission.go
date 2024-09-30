package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
)

const (
	ADMIN = "admin"
	USER  = "user"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, oke := c.Get(CurrentUser)
		if !oke {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		userRole := value.(models.User).Role

		hasRole := false
		if userRole == requiredRole {
			hasRole = true
		}
		if !hasRole || (userRole != ADMIN && userRole != USER) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "This endpoint is exclusive for role: " + requiredRole,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
