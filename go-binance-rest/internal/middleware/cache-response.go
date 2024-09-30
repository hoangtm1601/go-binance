package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hoangtm1601/go-binance-rest/internal/initializers"
	"github.com/rs/zerolog/log"
)

func CacheMiddleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip caching for non-GET requests
		if c.Request.Method != "GET" {
			c.Next()
			return
		}

		// Use the full request URL as the cache key
		key := c.Request.URL.String()

		// Try to get the cached response
		rdb := initializers.GetRedis()
		cachedResponse, err := rdb.Get(key).Bytes()
		if err == nil {
			// Cache hit: return the cached response
			c.Data(200, "application/json", cachedResponse)
			c.Abort() // Prevent further handlers from being called
			return
		}

		// Cache miss: capture the response
		customizedWriter := &responseWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = customizedWriter
		c.Next()

		// Cache the response if it's successful
		if c.Writer.Status() == 200 {
			go func() {
				err := rdb.Set(key, customizedWriter.body.Bytes(), duration).Err()
				if err != nil {
					log.Printf("Failed to set cache: %v", err)
				}
			}()
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
