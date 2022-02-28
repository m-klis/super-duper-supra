package service

import (
	"context"
	"noteapp/model"
	"noteapp/repository"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type NoteServiceImpl struct {
	NoteRepository repository.NoteRepository
	DB             *sqlx.DB
	Validate       *validator.Validate
}

func NewNoteService(noteRepo repository.NoteRepository, DB *sqlx.DB, validate *validator.Validate) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepo,
		DB:             DB,
		Validate:       validate,
	} // here is the magic
}

func (s *NoteServiceImpl) FindNotes(ctx context.Context) ([]model.NoteResponse, error) {
	db := s.DB

	notes, err := s.NoteRepository.FindNotes(ctx, db)
	if err != nil {
		return []model.NoteResponse{}, err
	}

	var noteResponses []model.NoteResponse
	for _, note := range notes {
		res := model.NoteResponse{
			ID:          note.ID,
			Title:       note.Title,
			Description: note.Description,
			Check:       note.Check,
			CreatedAt:   note.CreatedAt,
			UpdatedAt:   note.UpdatedAt,
		}
		noteResponses = append(noteResponses, res)
	}

	return noteResponses, nil
}

func (s *NoteServiceImpl) FindNote(ctx context.Context, noteId string) (model.NoteResponse, error) {
	db := s.DB

	notes, err := s.NoteRepository.FindNote(ctx, db, noteId)
	if err != nil {
		return model.NoteResponse{}, err
	}

	res := model.NoteResponse{
		ID:          notes.ID,
		Title:       notes.Title,
		Description: notes.Description,
		Check:       notes.Check,
		CreatedAt:   notes.CreatedAt,
		UpdatedAt:   notes.UpdatedAt,
	}

	return res, nil
}
