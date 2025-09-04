

        product-service/
        ├── controllers/
        │   └── product_controller.go
        ├── services/
        │   └── product_service.go
        ├── models/
        │   └── product.go
        ├── repository/
        │   └── product_repository.go
        ├── middleware/
        │   └── auth_middleware.go
        ├── config/
        │   └── config.go
        ├── routes/
        │   └── routes.go
        ├── main.go
        ├── go.mod
        ├── Dockerfile
        ├── .env
        ├── docker-compose.yml



        GET /api/products/
        GET /api/products/:id
        GET /api/products/search

        POST /api/products/
        PUT /api/products/:id
        DELETE /api/products/:id

        POST  /api/products/adjust-stock
