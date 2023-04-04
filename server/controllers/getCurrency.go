package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oaraujocesar/client-server-api/server/database"
	"github.com/oaraujocesar/client-server-api/server/utils"
)

func GetCurrenyHandler(w http.ResponseWriter, r *http.Request) {
	quote, err := utils.GetDollarQuote(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		w.Write([]byte("Request Timeout"))
	}

	json, err := json.Marshal(quote)
	if err != nil {
		log.Fatal(err)
	}

	err = database.InsertQuote(r.Context(), database.DB, quote)
	if err != nil {
		w.Write([]byte("Insertion Timeout"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
