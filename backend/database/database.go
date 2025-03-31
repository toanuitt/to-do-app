package database

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(){
    const(
        host = "localhost"
        port = 5432
        user = "postgres"
        password = "postgres"
        dbname = "postgres"
    )
    fmt.Println("Connecting to database...")
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    var err error
    DB, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal("loi roi",err)
    }
    err = DB.Ping()
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Successfully connected to the database")
    createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        priority INTEGER NOT NULL,
        category INTEGER NOT NULL,
        due_date TIMESTAMP,
        completed BOOLEAN NOT NULL DEFAULT FALSE
        );`

    _, err = DB.Exec(createTableSQL)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }
    log.Println("Table tasks created successfully")
}                                