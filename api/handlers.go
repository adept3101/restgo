package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

var user User

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{
		Message: "API is helthy",
		Status:  200,
	})
}

func createUser(db *sql.DB, name string) error {
	_, err := db.Exec(
		"INSERT INTO users(name) VALUES(?)",
		name,
	)
	return err
}

func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "BAD JSON", http.StatusBadRequest)
		return
	}

	// CreateUser(db, user.Name)
	if err := createUser(a.DB, user.Name); err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func selectUser(db *sql.DB, id uint8) (User, error) {
	err := db.QueryRow(
		"SELECT * FROM users WHERE id = ?",
		id,
	).Scan(user.ID, user.Name)
	return user, err

}

func selectUsers(db *sql.DB) error {
	_, err := db.Exec(
		"SELECT * FROM users;",
	)
	return err
}

func (a *App) GetUsr(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}


	if user, err := selectUser(a.DB, user.ID); user, err != nil {
		log.Println(err)
		http.Error(w, "Failed to get user", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad JSON", http.StatusBadRequest)
		return
	}

	if err := selectUsers(a.DB); err != nil {
		log.Println(err)
		http.Error(w, "Failed to get users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
