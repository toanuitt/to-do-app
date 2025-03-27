package routes

import (
	"github.com/gorilla/mux"
    "backend/controllers"
	"backend/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", listTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	// Task Management Routes
	router.HandleFunc("/tasks/{id}/complete", markTaskCompleted).Methods("PATCH")
	router.HandleFunc("/tasks/{id}/priority", setTaskPriority).Methods("PATCH")
	router.HandleFunc("/tasks/{id}/deadline", setTaskDeadline).Methods("PATCH")

	// Categorization Routes
	router.HandleFunc("/tasks/categories", listCategories).Methods("GET")
	router.HandleFunc("/tasks/{id}/category", assignCategory).Methods("PATCH")

	// Middleware
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthMiddleware)

	return router
}