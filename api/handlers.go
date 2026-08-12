package api

import (
	"database/sql"
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"strconv"
)

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

	if err != nil {
		return err
	}
	return nil
}

func (a *App) AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "BAD JSON", http.StatusBadRequest)
		return
	}

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
	var user User

	err := db.QueryRow(
		"SELECT id, name FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Name)

	return user, err
}

func (a *App) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 8)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := selectUser(a.DB, uint8(id))
	if err != nil {
		log.Println(err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func selectUsers(db *sql.DB) ([]User, error) {
	// var user User

	rows, err := db.Query(
		"SELECT * FROM users;",
	)
	usr := []User{}

	for rows.Next() {
		// u := usr
		var u User
		err := rows.Scan(&u.ID, &u.Name)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		defer rows.Close()
		usr = append(usr, u)
	}
	return usr, err
}

func (a *App) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	user, err := selectUsers(a.DB)
	if err != nil {
		log.Println(err)
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUser(db *sql.DB, id uint8) error {
	result, err := db.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (a *App) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 8)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	usr := deleteUser(a.DB, uint8(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usr)
}
