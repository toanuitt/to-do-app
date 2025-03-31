package controllers

import (
    "backend/database"
    "backend/models"
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `
        INSERT INTO tasks (title, description, priority, category, due_date, completed)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

    err := database.DB.QueryRow(
        query,
        task.Title,
        task.Description,
        task.Priority,
        task.Category,
        task.DueDate,
        task.Completed,
    ).Scan(&task.ID)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    query := `
        SELECT id, title, description, priority, category, due_date, completed 
        FROM tasks WHERE id = $1`

    err = database.DB.QueryRow(query, id).Scan(
        &task.ID,
        &task.Title,
        &task.Description,
        &task.Priority,
        &task.Category,
        &task.DueDate,
        &task.Completed,
    )

    if err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
    rows, err := database.DB.Query(`
        SELECT id, title, description, priority, category, due_date, completed 
        FROM tasks`)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var task models.Task
        err := rows.Scan(
            &task.ID,
            &task.Title,
            &task.Description,
            &task.Priority,
            &task.Category,
            &task.DueDate,
            &task.Completed,
        )
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tasks = append(tasks, task)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var task models.Task
    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    query := `
        UPDATE tasks 
        SET title = $1, description = $2, priority = $3, category = $4, due_date = $5, completed = $6
        WHERE id = $7
        RETURNING id`

    err = database.DB.QueryRow(
        query,
        task.Title,
        task.Description,
        task.Priority,
        task.Category,
        task.DueDate,
        task.Completed,
        id,
    ).Scan(&task.ID)

    if err != nil {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    result, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        http.Error(w, "Task not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}