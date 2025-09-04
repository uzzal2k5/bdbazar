package middleware

import (
	"net/http"
	"strings"
	"time"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware validates JWT, checks expiration, user status, and roles.
func RequireAuth(secret string, requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        	}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
		    fmt.Println("❌ Token parsing failed:", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// Check token expiration
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing exp in token"})
			return
		}

		// Optional: Check if user is active or blocked (if available)
		if blocked, ok := claims["is_blocked"].(bool); ok && blocked {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User is blocked"})
			return
		}

		if active, ok := claims["is_active"].(bool); ok && !active {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User account is inactive"})
			return
		}

		// Role check (if required)
		if len(requiredRoles) > 0 {
			rawRoles, ok := claims["roles"].([]interface{})
			if !ok {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid roles format in token"})
				return
			}

			hasRole := false
			for _, role := range rawRoles {
				roleStr, ok := role.(string)
				if !ok {
				    continue
				}
                for _, required := range requiredRoles {
					if roleStr == required {
						hasRole = true
						break

                    }
				}
                if hasRole {
                    break
                }
			}
			if !hasRole {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient role privileges"})
				return
			}
		}

        // Helper to safely extract string claims
		getStringClaim := func(key string) string {
			if val, ok := claims[key]; ok && val != nil {
				if s, ok := val.(string); ok {
					return s
				}
			}
			return ""
		}
		// Set user info in context
		c.Set("userID", claims["id"])
		c.Set("email", getStringClaim("email"))
		c.Set("mobile", getStringClaim("mobile"))
		c.Set("roles", claims["roles"])


		c.Next()
	}
}
