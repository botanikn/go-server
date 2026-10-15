package dataaccess

import (
	"context"

	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/storage"
)

type NotesRepository struct {
	storage *storage.Storage
}

func (n *NotesRepository) GetNoteRepository(id int) (entities.Note, error) {
	query := "SELECT * FROM notes WHERE id = $1"
	row := n.storage.Connection.QueryRow(context.Background(), query, id)
	var note entities.Note
	err := row.Scan(&note.ID, &note.Title, &note.Content)
	if err != nil {
		return entities.Note{}, err
	}
	return note, nil
}

func NewNoteRepository(strg *storage.Storage) NotesRepository {
	return NotesRepository{
		storage: strg,
	}
}

func (n *NotesRepository) FindAllNotes() ([]entities.Note, error) {
	query := "SELECT id, title, content FROM notes"
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
