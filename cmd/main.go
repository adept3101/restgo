package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"rest/api"
)

func main() {
	// path := os.Getenv("PATH")
	// fmt.Println("Path:", path)

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	app := &api.App{
		DB: db,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", app.AddUserHandler)
	// http.HandleFunc("GET /health", api.HealthHandler)
	mux.HandleFunc("GET /health", api.HealthHandler)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
