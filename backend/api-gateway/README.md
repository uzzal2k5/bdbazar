
        api-gateway/
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
        ├── .env
        ├── main.go
        ├── config/
        │   └── config.go
        ├── proxy/
        │   └── routes.go
        └── utils/
            └── logger.go



                             ┌────────────┐
                             │  Browser   │
                             │  / Mobile  │
                             └────┬───────┘
                                  │
                                  ▼
                         ┌─────────────────────┐
                         │     API Gateway     │
                         │  (Single Entrypoint)│
                         └────┬──────────────┬─┘
                              │              │
                     ┌────────▼────┐  ┌──────▼────────┐
                     │ auth-service│  │ product-service│
                     └─────────────┘  └────────────────┘



1. Create Service [actual port 8001 but used nodePort 8002]

curl -i -X POST http://localhost:8002/services \
  --data name=api-service \
  --data url=http://auth-service:8080

2. Create Route [actual port 8001 but used nodePort 8002]
curl -i -X POST http://localhost:8002/services/api-service/routes \
  --data paths[]=/api


2.1: Now call from Kong : http://localhost:8000/api/endpoint
2.2: Kong forward it to : http://auth-service:8080/endpoint

Here Kong Admin Port: 8001 use to register service


🔹 1. Register auth-service at /api/auth

        curl -i -X POST http://localhost:8001/services \
          --data name=auth-service \
          --data url=http://auth-service:3000

        curl -i -X POST http://localhost:8001/services/auth-service/routes \
          --data paths[]=/api/auth

🔹 2. Register product-service at /api/products

        curl -i -X POST http://localhost:8001/services \
          --data name=product-service \
          --data url=http://product-service:3001

        curl -i -X POST http://localhost:8001/services/product-service/routes \
          --data paths[]=/api/products


🔹 3. Register order-service at /api/orders

        curl -i -X POST http://localhost:8001/services \
          --data name=order-service \
          --data url=http://order-service:3002

        curl -i -X POST http://localhost:8001/services/order-service/routes \
          --data paths[]=/api/orders


🧪 Test Through Kong Proxy (port 8000)

        curl http://localhost:8000/api/auth/login
        curl http://localhost:8000/api/products/list
        curl http://localhost:8000/api/orders/create


These will internally route to:

        /api/auth → http://auth-service:3000
        /api/products → http://product-service:3001
        /api/orders → http://order-service:3002