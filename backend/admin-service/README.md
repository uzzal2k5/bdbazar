

            admin-service/
            ├── go.mod
            ├── go.sum
            ├── main.go                      # App entrypoint
            ├── Dockerfile
            ├── .env                         # Environment variables
            ├── config/
            │   └── config.go                # Loads environment, DB config
            ├── controllers/
            │   └── admin_controller.go      # All admin API handlers
            ├── middleware/
            │   └── auth_middleware.go       # JWT or session-based auth middleware
            ├── models/
            │   └── admin.go                 # Admin model
            │   └── activity_log.go          # Admin activity logs
            ├── repository/
            │   └── admin_repository.go      # Data access logic
            │   └── activity_log_repository.go
            ├── routes/
            │   └── routes.go                # Route definitions
            ├── services/
            │   └── admin_service.go         # Business logic layer
            │   └── activity_log_service.go
            ├── utils/
            │   └── logger.go                # Centralized logging
            │   └── response.go              # Common response formatting
            ├── seed/
            │   └── seed.go                  # Seed initial admin data
            └── README.md




🧩 Key Features

    - Admin Dashboard API
    - Approve/Block users
    - Approve/Block shops
    - Reset user passwords
    - Delete users
    - View platform metrics

    POST   /api/admins/spadm/login
    GET    /api/admins
    GET    /api/admins/:id
    POST   /api/admins
    PUT    /api/admins/:id
    DELETE /api/admins/:id
    GET    /api/admins/metrics
    GET    /api/admins/dashboard
    PATCH  /api/admins/user/:id/block
    PATCH  /api/admins/user/:id/approve
    POST   /api/admins/user/:id/reset-password
    DELETE /api/admins/user/:id
    PATCH  /api/admins/shop/:id/approve
    PATCH  /api/admins/shop/:id/block
    GET    /health
