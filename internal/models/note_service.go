package models

import (
	"context"

	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/storage"
)

type NoteService struct {
	storage *storage.Storage
}

func NewNoteService(strg *storage.Storage) NoteService {
	return NoteService{
		storage: strg,
	}
}

func (n *NoteService) GetAllNotes() ([]entities.Note, error) {
	query := "SELECT * FROM notes"
	rows, err := n.storage.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	notes := []entities.Note{}

	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (n *NoteService) GetNoteById(id int) (entities.Note, error) {
	query := "SELECT * FROM notes WHERE id = $1"

	row := n.storage.Connection.QueryRow(context.Background(), query, id)

	var note entities.Note
	if err := row.Scan(&note.ID, &note.Title, &note.Content); err != nil {
		return entities.Note{}, err
	}

	return note, nil
}

func (n *NoteService) CreateNote(note entities.Note) error {
	query := "INSERT INTO notes (title, content) VALUES ($1, $2)"
	_, err := n.storage.Connection.Exec(context.Background(), query, note.Title, note.Content)

	if err != nil {
		return err
	}

	return nil
}
