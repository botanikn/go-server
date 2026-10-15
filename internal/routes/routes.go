package routes

import (
	"log"
	"net/http"

	dataaccess "github.com/botanikn/go-server/internal/data_access"
	"github.com/botanikn/go-server/internal/handlers"
	"github.com/botanikn/go-server/internal/models"
	"github.com/botanikn/go-server/internal/storage"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	connString := "postgres://postgres:123456@0.0.0.0:5432/notes?sslmode=disable"
	storage, err := storage.NewStorage(connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	repository := dataaccess.NewNoteRepository(storage)
	noteService := models.NewNoteService(&repository)
	handler := handlers.NewHandler(noteService)

	mux.HandleFunc("GET /health", handler.HandleHealth)
	mux.HandleFunc("GET /notes", handler.HandleNotes)
	mux.HandleFunc("GET /notesXml", handler.HandleNotesXml)
	//mux.HandleFunc("POST /notes", handler.HandleCreateNotes)
	mux.HandleFunc("GET /notes/{id}", handler.HandleGetNoteByID)

	return mux
}
