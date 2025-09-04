package routes

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"

    "auth-service/controllers"
    "auth-service/middleware"
)

// AuthRoutes defines all API routes for the auth-service
func AuthRoutes(r *gin.Engine, authController controllers.AuthController,jwtSecret string) {
    // inside setup or router init
//     redisClient := redis.NewClient(&redis.Options{
//         Addr: "localhost:6379", // or use REDIS_URL
//     })
    // ───────────────────────────────
    // PUBLIC ROUTES
    // ───────────────────────────────
    // Public routes (no auth required)
    public := r.Group("/api/auth")
    {
        public.POST("/register", authController.Register)
        public.POST("/login", middleware.RateLimitMiddleware(), authController.Login)
        public.POST("/refresh", authController.Refresh)
        public.POST("/logout", authController.Logout)
    }

    // ───────────────────────────────
    // PROTECTED USER ROUTES
    // Protected user route (any role: buyer, seller, admin)
    // ───────────────────────────────
    protected := r.Group("/api/user")
    protected.Use(
        middleware.RequireAuth(jwtSecret, "buyer", "seller", "admin"),
        middleware.RateLimitMiddleware(),

    )
    protected.GET("/profile", func(c *gin.Context) {
            defer func() {
                if r := recover(); r != nil {
                    fmt.Println("🔥 Panic recovered in /profile:", r)
                    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
                }
            }()

            userID := c.MustGet("userID")
            email := c.MustGet("email")
            mobile := c.MustGet("mobile")
            roles := c.MustGet("roles")
            c.JSON(http.StatusOK, gin.H{
                "id":    userID,
                "email": email,
                "mobile": mobile,
                "roles": roles,
            })

    })


    // ───────────────────────────────
    // PROTECTED ADMIN ROUTES
    // ───────────────────────────────
    adminGroup := r.Group("/api/admin")
    adminGroup.Use(middleware.RequireAuth(jwtSecret, "admin"))
    {
        // Admin dashboard
        adminGroup.GET("/dashboard", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome Admin!"})
        })

        // Example: Delete user by ID
    	adminGroup.DELETE("/users/:id", func(c *gin.Context) {
    		userID := c.Param("id")
    		// Placeholder for delete logic: deleteUserByID(userID)
    		c.JSON(http.StatusOK, gin.H{
    			"message": "User deleted successfully",
    			"user_id": userID,
    		})
    	})

        // Example: System status check
    	adminGroup.GET("/status", func(c *gin.Context) {
    		c.JSON(http.StatusOK, gin.H{
    			"status":  "OK",
    			"uptime":  "123h",
    			"version": "v1.0.0",
    		})
    	})

        // Example: System status check
        adminGroup.GET("/settings", func(c *gin.Context) {
            // Example settings — these could be loaded from env, a database, or config service
            settings := gin.H{
                "app_name":       "BDBazar Marketplace",
                "version":        "v1.0.0",
                "maintenance":    false,
                "supported_roles": []string{"buyer", "seller", "admin"},
                "max_login_attempts": 5,
                "token_expiry_minutes": 1440,
            }

            c.JSON(http.StatusOK, gin.H{
                "settings": settings,
            })
        })

    }

    // ───────────────────────────────
    // PROTECTED SELLER ROUTES
    // Seller-only routes
    // ───────────────────────────────
    sellerGroup := r.Group("/api/seller")
    sellerGroup.Use(middleware.RequireAuth(jwtSecret, "seller"))
    {
        sellerGroup.GET("/dashboard", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{"message": "Welcome Seller!"})
        })
    }
}
