

        auth-service/
        ├── cmd/                # Entry point for the application
        │   └── main.go
        ├── config/             # Loads environment variables using Viper
        │   └── config.go
        ├── handlers/           # HTTP handlers (controllers)
        │   └── auth_handler.go
        ├── middleware/         # Middleware (JWT auth, logging) [Empty for now]
        ├── models/             # Database models [To be added]
        ├── repository/         # Database interaction layer [To be added]
        ├── routes/             # HTTP route setup
        │   └── routes.go
        ├── services/           # Business logic layer [To be added]
        ├── utils/              # Utility functions [Optional]
        ├── .env                # Environment variables
        └── go.mod / go.sum     # Go module files


Endpoints Available:

    - POST /api/auth/register

    - POST /api/auth/login