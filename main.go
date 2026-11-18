package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
)

//go:embed quotes.json
var quotesFile []byte

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source"`
	Year   string `json:"year"`
}

type quoteData struct {
	Quotes []Quote `json:"quotes"`
}

var quotes []Quote

func init() {
	var data quoteData
	if err := json.Unmarshal(quotesFile, &data); err != nil {
		log.Fatalf("failed to parse embedded quotes: %v", err)
	}
	quotes = data.Quotes
}

func randomQuote() Quote {
	return quotes[rand.Intn(len(quotes))]
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	q := randomQuote()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(q.Text + " - " + q.Author))
}

func quoteJSONHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(randomQuote())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /getquotejson", quoteJSONHandler)

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
