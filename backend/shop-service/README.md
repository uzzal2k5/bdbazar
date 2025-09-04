

        shop-service/
        ├── controllers/
        │   └── shop_controller.go
        ├── models/
        │   └── shop.go
        ├── repository/
        │   └── shop_repository.go
        ├── services/
        │   └── shop_service.go
        ├── routes/
        │   └── routes.go
        ├── middleware/
        │   └── auth_middleware.go
        ├── config/
        │   └── config.go
        ├── cmd/main.go




        #API List
        GET /api/shops
        GET /api/shops/search
        GET /api/shops/:id
        GET /api/shops/dashboard

        POST /api/shops/

        PUT /api/shops/:id
        DELETE /api/shops/:id