package routes

import (
	"log"
	"net/http"

	"github.com/botanikn/go-server/internal/handlers"
	"github.com/botanikn/go-server/internal/storage"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	connString := "postgres://postgres:123456@0.0.0.0:5432/notes?sslmode=disable"
	storage, err := storage.NewStorage(connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	handler := &handlers.Handler{
		Storage: storage,
	}
	mux.HandleFunc("POST /health", handler.HandleHealth)
	mux.HandleFunc("GET /notes", handler.HandleNotes)
	mux.HandleFunc("POST /notes", handler.HandleCreateNotes)
	mux.HandleFunc("GET /notes/{id}", handler.HandleGetNoteByID)

	return mux
}
