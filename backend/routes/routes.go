package routes

import (
	"github.com/gorilla/mux"
    "backend/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", controllers.ListTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// Task Management Routes
	router.HandleFunc("/api/tasks/{id}/complete", controllers.MarkTaskCompleted).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}/priority", controllers.SetTaskPriority).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}/deadline", controllers.SetTaskDeadline).Methods("PATCH")

	// Categorization Routes
	router.HandleFunc("/api/tasks/categories", controllers.ListCategories).Methods("GET")
	router.HandleFunc("/api/tasks/{id}/category", controllers.AssignCategory).Methods("PATCH")

	return router
}