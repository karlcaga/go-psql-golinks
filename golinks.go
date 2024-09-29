package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /{shortlink}", handleShortLink)
	http.HandleFunc("GET /{shortlink}/", handleShortLink)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShortLink(w http.ResponseWriter, r *http.Request) {
	shortlink := r.PathValue("shortlink")
	fmt.Fprintf(w, "%v", shortlink)
}
