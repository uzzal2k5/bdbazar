package config

const DefaultPort = "8000"

var ServiceRoutes = map[string]string{
    "auth":     "http://auth-service:8080",
    "shop":     "http://shop-service:8084",
    "product":  "http://product-service:8085",
    "order":    "http://order-service:8086",
    "payment-": "http://payment-service:8087",
    "shipping": "http://shipping-service:8088",
    "admin":    "http://admin-service:8090",
}
