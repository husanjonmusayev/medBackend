package database

import (
	"database/sql"
	"log"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
