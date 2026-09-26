package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/botanikn/go-server/internal/routes"
)

func main() {
	mux := routes.RegisterRoutes()
	fmt.Println("Server is starting")

	server := http.Server{
		Addr:    "127.0.0.1:8082",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server has fallen because of error: %v", err)
	}
}
