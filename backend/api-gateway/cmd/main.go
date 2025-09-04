package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "api-gateway/proxy"
    "api-gateway/config"
    "api-gateway/utils"
    "github.com/joho/godotenv"
)


func main() {
    // Load env variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("⚠️  No .env file found, using system env variables")
    }
    router := mux.NewRouter()

    proxy.RegisterRoutes(router)

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("BDBAZAR API Gateway is running"))
    })

    port := os.Getenv("GATEWAY_PORT")
    if port == "" {
        port = config.DefaultPort
    }

    utils.LogInfo("Starting API Gateway on port " + port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
