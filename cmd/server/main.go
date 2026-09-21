package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is alive and got "+"request")
}

func main() {
	http.HandleFunc("POST /health", handleHealth)

	fmt.Println("Server is starting")

	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		log.Fatalf("application exit")
	}
}
