package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	contentLength := header.Get("Content-Length")
	userAgent := header.Get("User-Agent")
	contentType := header.Get("Content-Type")

	if contentType != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,
		"Server is alive and has received request with contentLength: '%v' and userAgent: '%v' and contentType: '%v'",
		contentLength,
		userAgent,
		contentType,
	)

	fmt.Fprintf(os.Stdout,
		"Server is alive and has received request with contentLength: '%v' and userAgent: '%v' and contentType: '%v'",
		contentLength,
		userAgent,
		contentType,
	)
}

func main() {
	http.HandleFunc("POST /health", handleHealth)

	fmt.Println("Server is starting")

	err := http.ListenAndServe("127.0.0.1:8082", nil)
	if err != nil {
		log.Fatalf("Server has fallen because of error: %v", err)
	}
}
