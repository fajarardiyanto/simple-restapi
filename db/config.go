package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Config() {
	var err error
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("Can't load .env file %s", err)
	}

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, _ := sql.Open("postgres", dbInfo)
	err = db.Ping()
	if err != nil {
		fmt.Printf("Can't load databse %s", err)
	}

	DB = db

	TodoSchema()
}
