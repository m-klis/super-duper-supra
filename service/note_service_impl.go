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

func (s *NoteServiceImpl) FindAll(ctx context.Context) []model.NoteResponse {
	db := s.DB

	notes := s.NoteRepository.FindAll(ctx, db)

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

	return noteResponses
}
