package proxy

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "github.com/gorilla/mux"
    "api-gateway/config"
)

func createReverseProxy(target string) *httputil.ReverseProxy {
    url, err := url.Parse(target)
    if err != nil {
        panic("Invalid service URL: " + target)
    }
    return httputil.NewSingleHostReverseProxy(url)
}
// api-gateway-url/api/auth -> <upstream> -> microservice-base-url/api/auth
// base_url/api/products
// base_url/api/shop
func RegisterRoutes(r *mux.Router) {
    r.PathPrefix("/api/auth/").Handler(http.StripPrefix("/auth", createReverseProxy(config.ServiceRoutes["auth"])))
    r.PathPrefix("/shop/").Handler(http.StripPrefix("/shop", createReverseProxy(config.ServiceRoutes["shop"])))
    r.PathPrefix("/products/").Handler(http.StripPrefix("/products", createReverseProxy(config.ServiceRoutes["product"])))
    r.PathPrefix("/orders/").Handler(http.StripPrefix("/orders", createReverseProxy(config.ServiceRoutes["order"])))
    r.PathPrefix("/payments/").Handler(http.StripPrefix("/payments", createReverseProxy(config.ServiceRoutes["payment"])))
    r.PathPrefix("/shipping/").Handler(http.StripPrefix("/shipping", createReverseProxy(config.ServiceRoutes["shipping"])))
    r.PathPrefix("/admin/").Handler(http.StripPrefix("/admin", createReverseProxy(config.ServiceRoutes["admin"])))
}
