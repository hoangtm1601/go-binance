package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/models/dto"
	"net/http"
)

const (
	Pagination string = "pagination"
)

func BindPagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		pagination := &dto.PaginationDto{
			Page:    1,
			PerPage: 20,
		}
		if err := c.ShouldBindQuery(pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			c.Abort()
			return
		}
		c.Set(Pagination, pagination)
		c.Next()
	}
}
