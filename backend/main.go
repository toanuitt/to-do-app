package main

import (
    "backend/database"
    "backend/routes"
    "log"
    "net/http"
)

func main() {
    // Initialize database connection
    database.InitDB()
    defer database.DB.Close()

    // Initialize router
    router := routes.Router()

    // Add CORS middleware
    handler := corsMiddleware(router)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}