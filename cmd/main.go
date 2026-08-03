package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"rest/api"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	app := &api.App{
		DB: db,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("POST /user", app.AddUserHandler)
	mux.HandleFunc("GET /health", api.HealthHandler)
	mux.HandleFunc("GET /users", app.GetUsers)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
