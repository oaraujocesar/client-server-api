package main

import (
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/oaraujocesar/client-server-api/server/controllers"
	"github.com/oaraujocesar/client-server-api/server/database"
)

func main() {
	database.Connect()
	defer database.DB.Close()

	http.HandleFunc("/cotacao", controllers.GetCurrenyHandler)

	fmt.Printf("Server is running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
