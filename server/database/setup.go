package database

import "database/sql"

var DB *sql.DB

func Connect() {
	db, err := sql.Open("sqlite3", "quotes.db")
	if err != nil {
		panic("DATABASE: Failed to connect.")
	}

	DB = db
}
