package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Logger is a middleware that logs the incoming HTTP requests.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf(
			"%s %s %s %s",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			time.Since(start),
		)
	}
}

// CORS is a middleware that handles Cross-Origin Resource Sharing.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// RequestID is a middleware that injects a request ID into the context.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}

// RateLimiter is a middleware that limits the rate of incoming requests.
func RateLimiter(limit int) gin.HandlerFunc {
	tokens := make(chan struct{}, limit)

	for i := 0; i < limit; i++ {
		tokens <- struct{}{}
	}

	return func(c *gin.Context) {
		select {
		case <-tokens:
			c.Next()
			tokens <- struct{}{}
		default:
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
	}
}
