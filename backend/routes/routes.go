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

	return router
}