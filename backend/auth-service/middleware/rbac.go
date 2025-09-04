package middleware

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

// __define-ocg__: Middleware for RBAC with multiple roles
func RequireRoles(secret string, allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.GetHeader("Authorization")
        if auth == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            return
        }

        parts := strings.Split(auth, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
            return
        }

        tokenStr := parts[1]
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }

        rolesIface, exists := claims["roles"]
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "No roles in token"})
            return
        }

        roles := make([]string, 0)
        switch v := rolesIface.(type) {
        case []interface{}:
            for _, r := range v {
                if s, ok := r.(string); ok {
                    roles = append(roles, s)
                }
            }
        case []string:
            roles = v
        }

        for _, role := range roles {
            for _, allowed := range allowedRoles {
                if role == allowed {
                    c.Set("user_id", claims["id"])
                    c.Set("user_email", claims["email"])
                    c.Set("user_roles", roles)
                    c.Next()
                    return
                }
            }
        }

        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
    }
}
