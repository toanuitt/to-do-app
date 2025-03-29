package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Kết nối đến PostgreSQL trong Docker
	const (
		host     = "postgres" // Dùng tên service trong Docker Compose
		port     = 5432       // Port mặc định của PostgreSQL
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
	)

	// Chuỗi kết nối PostgreSQL
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	// Mở kết nối đến PostgreSQL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Lỗi khi mở kết nối đến database:", err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	if err = db.Ping(); err != nil {
		log.Fatal("Không thể kết nối đến PostgreSQL:", err)
	}

	fmt.Println("✅ Kết nối thành công đến PostgreSQL trong Docker!")
}
