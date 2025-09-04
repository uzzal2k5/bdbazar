// middleware/rate_limit.go
package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitor struct {
	Attempts int
	LastSeen time.Time
}

var visitors = make(map[string]*visitor)
var mu sync.Mutex

const (
	LimitAttempts = 5
	Window        = 15 * time.Minute
)

func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.LastSeen) > Window {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

func init() {
	go cleanupVisitors()
}

// RateLimitMiddleware limits requests from the same IP
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		v, exists := visitors[ip]
		if !exists {
			v = &visitor{Attempts: 1, LastSeen: time.Now()}
			visitors[ip] = v
		} else {
			v.Attempts++
			v.LastSeen = time.Now()
		}
		mu.Unlock()

		if v.Attempts > LimitAttempts {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many login attempts. Please try again later.",
			})
			return
		}

		c.Next()
	}
}
