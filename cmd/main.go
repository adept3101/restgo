package main

import (
	"fmt"
	"net/http"
	"rest/api"
)

func main() {
	db := api.ConnDB()

	defer db.Close()

	app := &api.App{
		DB: db,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.HealthHandler)
	mux.HandleFunc("POST /user", app.AddUserHandler)
	mux.HandleFunc("GET /user/{id}", app.GetUserHandler)
	mux.HandleFunc("GET /users", app.GetUsersHandler)
	mux.HandleFunc("DELETE /delete/{id}", app.DeleteUserHandler)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe("localhost:8080", mux)
}
