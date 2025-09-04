

        payment-service/
        ├── main.go
        ├── go.mod
        ├── config/
        │   └── config.go
        ├── models/
        │   └── payment.go
        ├── repository/
        │   └── payment_repository.go
        ├── services/
        │   └── payment_service.go
        ├── controllers/
        │   └── payment_controller.go
        ├── middleware/
        │   └── auth_middleware.go
        ├── routes/
        │   └── routes.go
        ├── Dockerfile
        ├── .env

        POST   /api/payments/
        GET    /api/payments/buyer
        GET    /api/payments/seller
        POST   /api/payments/:id/complete
