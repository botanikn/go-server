package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/botanikn/go-server/internal/application/service"
	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/presentation/views"
)

type NoteHandler struct {
	noteService service.NoteService
}

func NewNoteHandler(noteService service.NoteService) NoteHandler {
	return NoteHandler{
		noteService: noteService,
	}
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

	result, err := h.noteService.GetAllNotes()

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

	result, err := h.noteService.GetAllNotes()

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

func (h *NoteHandler) HandleGetNoteByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	form := service.FindNoteByIdForm{
		Id: id,
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
	var note entities.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := h.noteService.CreateNote(note)

	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Note created")
}
