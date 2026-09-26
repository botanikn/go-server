package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/storage"
)

type Handler struct {
	Storage *storage.Storage
}

func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleNotes(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM notes"
	rows, err := h.Storage.Connection.Query(context.Background(), query)
	if err != nil {
		http.Error(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}

	notes := []entities.Note{}

	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			http.Error(w, "Failed to scan note", http.StatusInternalServerError)
			return
		}
		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "Failed to encode notes", http.StatusInternalServerError)
	}
}

func (h *Handler) HandleGetNoteByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM notes WHERE id = $1"

	row := h.Storage.Connection.QueryRow(context.Background(), query, id)

	var note entities.Note
	if err := row.Scan(&note.ID, &note.Title, &note.Content); err != nil {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "Failed to encode note", http.StatusInternalServerError)
	}

}

func (h *Handler) HandleCreateNotes(w http.ResponseWriter, r *http.Request) {
	var note entities.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO notes (title, content) VALUES ($1, $2)"
	_, err := h.Storage.Connection.Exec(context.Background(), query, note.Title, note.Content)

	if err != nil {
		http.Error(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Note created")
}
