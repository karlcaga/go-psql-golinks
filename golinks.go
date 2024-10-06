package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	log.SetPrefix("golinks: ")
	if os.Getenv("CONN_STR") == "" {
		envErr := godotenv.Load()
		if envErr != nil {
			log.Fatal("Error loading .env file")
		}
	}

	var dbErr error
	db, dbErr = sql.Open("postgres", os.Getenv("CONN_STR"))
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	http.HandleFunc("GET /{shortlink}", handleShortLink)
	http.HandleFunc("GET /{shortlink}/", handleShortLink)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShortLink(w http.ResponseWriter, r *http.Request) {
	shortlink := r.PathValue("shortlink")
	url, err := getURL(shortlink)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error obtaining URL: %v", err), 404)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func getURL(shortlink string) (string, error) {
	var url string
	if err := db.QueryRow("SELECT url FROM links WHERE shortlink = $1", shortlink).Scan(&url); err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("shortlink %v: unknown URL", shortlink)
		}
		return "", fmt.Errorf("shortlink %v: %v ", shortlink, err)
	}
	return url, nil
}
