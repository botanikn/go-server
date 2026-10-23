package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/botanikn/go-server/internal/application/errors"
	"github.com/botanikn/go-server/internal/application/service"
	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/presentation/views"
)

type NoteHandler struct {
	noteService service.NoteService
	authService service.AuthService
}

func NewNoteHandler(noteService service.NoteService, authService service.AuthService) NoteHandler {
	return NoteHandler{
		noteService: noteService,
		authService: authService,
	}
}

type NotesResponse struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *NoteHandler) HandleHealth(w http.ResponseWriter, r *http.Request) {
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

func (h *NoteHandler) HandleNotes(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err := h.authService.ValidateToken(token)

	if err != nil {
		if err.Error() == "token has expired" {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		return
	}

	result, err := h.noteService.GetAllNotes(userId)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	formatedResult, err := views.FormatJson(result)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", formatedResult.ContentType)
	w.WriteHeader(http.StatusOK)

	w.Write(formatedResult.Body)
}

func (h *NoteHandler) HandleNotesXml(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err := h.authService.ValidateToken(token)

	if err != nil {
		if err.Error() == "token has expired" {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		return
	}

	result, err := h.noteService.GetAllNotes(userId)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	formatedResult, err := views.FormatXml(result)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", formatedResult.ContentType)
	w.WriteHeader(http.StatusOK)

	w.Write(formatedResult.Body)
}

func (h *NoteHandler) HandleNotesByUserId(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err := h.authService.ValidateToken(token)

	if err != nil {
		if err.Error() == "token has expired" {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		return
	}

	result, err := h.noteService.GetUsersNotes(userId)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	response := []NotesResponse{}

	for _, note := range result {
		response = append(response, NotesResponse{
			Title:   note.Title,
			Content: note.Content,
		})
	}

	formatedResult, err := views.FormatJson(response)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", formatedResult.ContentType)
	w.WriteHeader(http.StatusOK)

	w.Write(formatedResult.Body)
}

func (h *NoteHandler) HandleGetNoteByID(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}

	userId, err := h.authService.ValidateToken(token)

	if err != nil {
		if err.Error() == "token has expired" {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	form := service.FindNoteByIdForm{
		NoteId: id,
		UserId: userId,
	}

	result, err := h.noteService.GetNoteById(form)

	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusInternalServerError)
		return
	}

	formatedResult, err := views.FormatJson(result)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", formatedResult.ContentType)
	w.WriteHeader(http.StatusOK)

	w.Write(formatedResult.Body)

}

func (h *NoteHandler) HandleCreateNotes(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		http.Error(w, "Missing token", http.StatusUnauthorized)
		return
	}
	userId, err := h.authService.ValidateToken(token)

	if err != nil {
		if err.Error() == errors.ErrExpiredToken {
			http.Error(w, "Token has expired", http.StatusUnauthorized)
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		}
		return
	}

	var note entities.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	result := h.noteService.CreateNote(note, userId)

	if result != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Note created")
}
