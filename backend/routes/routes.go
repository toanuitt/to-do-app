package routes

import (
	"github.com/gorilla/mux"
    "backend/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", controllers.ListTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

	// Task Management Routes
	router.HandleFunc("/tasks/{id}/complete", controllers.MarkTaskCompleted).Methods("PATCH")
	router.HandleFunc("/tasks/{id}/priority", controllers.SetTaskPriority).Methods("PATCH")
	router.HandleFunc("/tasks/{id}/deadline", controllers.SetTaskDeadline).Methods("PATCH")

	// Categorization Routes
	router.HandleFunc("/tasks/categories", controllers.ListCategories).Methods("GET")
	router.HandleFunc("/tasks/{id}/category", controllers.AssignCategory).Methods("PATCH")

	return router
}