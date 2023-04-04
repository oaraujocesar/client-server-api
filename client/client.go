package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/oaraujocesar/client-server-api/server/types"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var quote types.Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("response.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, fmt.Sprintf("Dólar: %s", quote.Data.Bid))
	if err != nil {
		panic(err)
	}

}
