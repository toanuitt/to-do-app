package main

import (
    "backend/database"
    "backend/routes"
    "log"
    "net/http"
    "path/filepath"
)

func main() {
    // Initialize database connection
    database.InitDB()
    defer database.DB.Close()

    // Initialize router
    router := routes.Router()


    // apiRouter := router.PathPrefix("/api").Subrouter()
    // routes.RegisterAPIRoutes(apiRouter)
    // Serve static files from frontend directory
    frontendDir := filepath.Join("..", "frontend")
    fs := http.FileServer(http.Dir(frontendDir))
    router.PathPrefix("/").Handler(http.StripPrefix("/", fs))

    // Add CORS middleware
    handler := corsMiddleware(router)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}