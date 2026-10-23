package repository

import (
	"context"

	"github.com/botanikn/go-server/internal/entities"
	"github.com/botanikn/go-server/internal/storage"
)

type NotesRepository struct {
	storage *storage.Storage
}

func NewNoteRepository(strg *storage.Storage) NotesRepository {
	return NotesRepository{
		storage: strg,
	}
}

func (n *NotesRepository) GetNoteById(id int) (entities.Note, error) {
	query := "SELECT id, title, content, user_id FROM notes WHERE id = $1"
	row := n.storage.Connection.QueryRow(context.Background(), query, id)

	var note entities.Note
	err := row.Scan(&note.ID, &note.Title, &note.Content, &note.UserId)
	if err != nil {
		return entities.Note{}, err
	}

	return note, nil
}

func (n *NotesRepository) FindAllNotes() ([]entities.Note, error) {
	query := "SELECT id, title, content, user_id FROM notes"
	rows, err := n.storage.Connection.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	notes := []entities.Note{}

	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.UserId); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (n *NotesRepository) GetNotesByUserId(userId int) ([]entities.Note, error) {
	query := "SELECT title, content FROM notes WHERE user_id = $1"
	rows, err := n.storage.Connection.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	notes := []entities.Note{}

	for rows.Next() {
		var note entities.Note
		if err := rows.Scan(&note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}

func (n *NotesRepository) CreateNote(note entities.Note, userId int) error {
	query := "INSERT INTO notes (title, content, user_id) VALUES ($1, $2, $3)"
	_, err := n.storage.Connection.Exec(context.Background(), query, note.Title, note.Content, userId)

	if err != nil {
		return err
	}

	return nil
}
