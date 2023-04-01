package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("sqlite3", "quotes.db")
	if err != nil {
		panic("DATABASE: Failed to connect.")
	}

	DB = db

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY,
            bid REAL
		)
	`)
	if err != nil {
		log.Fatal(err)
		return
	}
}
