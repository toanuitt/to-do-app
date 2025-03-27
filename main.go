package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Task struct represents a todo task
type Task struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

// TaskManager manages the tasks
type TaskManager struct {
	mu       sync.RWMutex
	tasks    map[int]Task
	nextID   int
}

// NewTaskManager creates a new TaskManager
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

// CreateTask adds a new task
func (tm *TaskManager) CreateTask(text string) Task {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := Task{
		ID:        tm.nextID,
		Text:      text,
		Completed: false,
	}
	tm.tasks[tm.nextID] = task
	tm.nextID++
	return task
}

// GetTasks returns all tasks
func (tm *TaskManager) GetTasks() []Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	tasks := make([]Task, 0, len(tm.tasks))
	for _, task := range tm.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// UpdateTask updates a task's text
func (tm *TaskManager) UpdateTask(id int, text string) (Task, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return Task{}, false
	}
	task.Text = text
	tm.tasks[id] = task
	return task, true
}

// DeleteTask removes a task
func (tm *TaskManager) DeleteTask(id int) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	_, exists := tm.tasks[id]
	if !exists {
		return false
	}
	delete(tm.tasks, id)
	return true
}

// ToggleTask toggles task completion status
func (tm *TaskManager) ToggleTask(id int, completed bool) (Task, bool) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task, exists := tm.tasks[id]
	if !exists {
		return Task{}, false
	}
	task.Completed = completed
	tm.tasks[id] = task
	return task, true
}

// Global task manager
var taskManager = NewTaskManager()

func main() {
	r := mux.NewRouter()

	// Create a new task
	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		var taskInput struct {
			Text string `json:"text"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task := taskManager.CreateTask(taskInput.Text)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}).Methods("POST")

	// Get all tasks
	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := taskManager.GetTasks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}).Methods("GET")

	// Update a task
	r.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		var taskInput struct {
			Text string `json:"text"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&taskInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, ok := taskManager.UpdateTask(id, taskInput.Text)
		if !ok {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}).Methods("PUT")

	// Delete a task
	r.HandleFunc("/tasks/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		if !taskManager.DeleteTask(id) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}).Methods("DELETE")

	// Toggle task completion
	r.HandleFunc("/tasks/{id}/toggle", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		var toggleInput struct {
			Completed bool `json:"completed"`
		}
		
		if err := json.NewDecoder(r.Body).Decode(&toggleInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task, ok := taskManager.ToggleTask(id, toggleInput.Completed)
		if !ok {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}).Methods("PATCH")

	// Serve the HTML file
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}