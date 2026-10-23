package service

import (
	"fmt"

	"github.com/botanikn/go-server/internal/application/errors"
	"github.com/botanikn/go-server/internal/entities"
)

type NoteRepositoryInterface interface {
	GetNoteById(id int) (entities.Note, error)
	FindAllNotes() ([]entities.Note, error)
	CreateNote(note entities.Note, userId int) error
	GetNotesByUserId(userId int) ([]entities.Note, error)
}

type FindNoteByIdForm struct {
	NoteId int
	UserId int
}

type NoteService struct {
	noteRep NoteRepositoryInterface
	authRep AuthRepositoryInterface
}

func NewNoteService(noteRep NoteRepositoryInterface, authRep AuthRepositoryInterface) NoteService {
	return NoteService{
		noteRep: noteRep,
		authRep: authRep,
	}
}

func (n *NoteService) GetAllNotes(userId int) ([]entities.Note, error) {
	role, err := n.authRep.GetUserRoleByUserId(userId)

	if err != nil {
		return nil, err
	}

	if role != "admin" {
		return nil, fmt.Errorf(errors.ErrAccessDenied)
	}
	notes, err := n.noteRep.FindAllNotes()

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NoteService) GetUsersNotes(userId int) ([]entities.Note, error) {
	notes, err := n.noteRep.GetNotesByUserId(userId)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (n *NoteService) GetNoteById(form FindNoteByIdForm) (entities.Note, error) {
	result, err := n.noteRep.GetNoteById(form.NoteId)

	if result.UserId != form.UserId {
		return entities.Note{}, fmt.Errorf(errors.ErrAccessDenied)
	}

	if err != nil {
		return entities.Note{}, err
	}

	return result, nil
}

func (n *NoteService) CreateNote(note entities.Note, userId int) error {
	return n.noteRep.CreateNote(note, userId)
}
