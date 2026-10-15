package models

import (
	"github.com/botanikn/go-server/internal/entities"
)

type NoteRepositoryInterface interface {
	GetNoteRepository(id int) (entities.Note, error)
	FindAllNotes() ([]entities.Note, error)
}

type FindNoteByIdForm struct {
	Id int
}

type NoteService struct {
	rep NoteRepositoryInterface
}

func NewNoteService(rep NoteRepositoryInterface) NoteService {
	return NoteService{
		rep: rep,
	}
}

func (n *NoteService) GetAllNotes() ([]entities.Note, error) {
	notes, err := n.rep.FindAllNotes()

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NoteService) GetNoteById(form FindNoteByIdForm) (entities.Note, error) {
	result, err := n.rep.GetNoteRepository(form.Id)
	return result, err
}

// func (n *NoteService) CreateNote(note entities.Note) error {
// 	query := "INSERT INTO notes (title, content) VALUES ($1, $2)"
// 	_, err := n.storage.Connection.Exec(context.Background(), query, note.Title, note.Content)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
