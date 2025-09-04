package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"context"
)

var ctx = context.Background()

func LoginRateLimiter(rdb *redis.Client, maxAttempts int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := fmt.Sprintf("login_attempts:%s", clientIP)

		// Get current attempt count
		attempts, err := rdb.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Redis error"})
			c.Abort()
			return
		}

		if attempts >= maxAttempts {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many login attempts. Please try again later.",
			})
			c.Abort()
			return
		}

		// Increment login attempts
		if err := rdb.Incr(ctx, key).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record attempt"})
			c.Abort()
			return
		}

		// Set expiration on first attempt
		if attempts == 0 {
			rdb.Expire(ctx, key, duration)
		}

		c.Next()
	}
}
