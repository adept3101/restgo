package api

import (
	"database/sql"
	// "fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	DB *sql.DB
}

func ConnDB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error load .env:", err)
	}

	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return  db
}
