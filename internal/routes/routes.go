package routes

import (
	"log"
	"net/http"

	"github.com/botanikn/go-server/internal/application/service"
	"github.com/botanikn/go-server/internal/infrastructure/repository"
	cryptoService "github.com/botanikn/go-server/internal/infrastructure/service"
	"github.com/botanikn/go-server/internal/presentation/handlers"
	"github.com/botanikn/go-server/internal/storage"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	connString := "postgres://postgres:123456@0.0.0.0:5432/notes?sslmode=disable"
	storage, err := storage.NewStorage(connString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Infrastructure layer
	noteRepository := repository.NewNoteRepository(storage)

	// Application layer
	noteService := service.NewNoteService(&noteRepository)

	// Presentation layer
	handler := handlers.NewNoteHandler(noteService)

	// Infrastructure layer for auth
	authRepository := repository.NewAuthRepository(storage)
	cryptoService := cryptoService.NewCryptoService()

	// Application layer for auth
	authService := service.NewAuthService(authRepository, cryptoService)

	// Presentation layer for auth
	authHandler := handlers.NewAuthHandler(authService)

	// Auth routes
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	// Note routes
	mux.HandleFunc("GET /health", handler.HandleHealth)
	mux.HandleFunc("GET /notes", handler.HandleNotes)
	mux.HandleFunc("GET /notesXml", handler.HandleNotesXml)
	mux.HandleFunc("POST /notes", handler.HandleCreateNotes)
	mux.HandleFunc("GET /notes/{id}", handler.HandleGetNoteByID)

	return mux
}
