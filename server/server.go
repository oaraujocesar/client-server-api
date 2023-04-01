package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oaraujocesar/client-server-api/server/database"
	"github.com/oaraujocesar/client-server-api/server/utils"
)

func main() {
	database.Connect()
	defer database.DB.Close()

	_, err := database.DB.Exec(`
		CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY,
            bid REAL
		)
	`)
	if err != nil {
		log.Fatal(err)
		return
	}

	http.HandleFunc("/cotacao", getCurrenyHandler)

	http.ListenAndServe(":8080", nil)
}

func getCurrenyHandler(w http.ResponseWriter, r *http.Request) {
	quote, err := utils.GetDollarQuote(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		log.Fatal(err)
		w.Write([]byte("Request Timeout"))
	}

	json, err := json.Marshal(quote)
	if err != nil {
		log.Fatal(err)
	}

	err = database.InsertQuote(r.Context(), database.DB, quote)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Insertion Timeout"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
