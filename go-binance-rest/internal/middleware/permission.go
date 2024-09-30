package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/models"
)

const (
	ADMIN = "admin"
	USER  = "user"
)

func RequireRole(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, ok := c.Get(CurrentUser)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		userRole := value.(models.User).Role

		hasRole := false
		for _, role := range requiredRoles {
			if userRole == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied. Required roles: " + strings.Join(requiredRoles, ", "),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
