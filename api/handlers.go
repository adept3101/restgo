package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message: "API is helthy",
		Status:  200,
	})
}

func CreateUser(db *sql.DB, name string) error {
	_, err := db.Exec(
		"INSERT INTO users(name) VALUES(?)",
		name,
	)
	return err
}

func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "BAD JSON", http.StatusBadRequest)
		return
	}

	// CreateUser(db, user.Name)
	if err := CreateUser(a.DB, user.Name); err != nil {
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
