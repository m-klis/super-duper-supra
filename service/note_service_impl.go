package service

import (
	"context"
	"noteapp/entity"
	"noteapp/exception"
	"noteapp/model"
	"noteapp/repository"
	"strconv"
	"time"

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

func (s *NoteServiceImpl) FindNotes(ctx context.Context) ([]model.NoteResponseDateString, error) {
	db := s.DB

	notes, err := s.NoteRepository.FindNotes(ctx, db)
	if err != nil {
		return []model.NoteResponseDateString{}, err
	}

	var noteResponses []model.NoteResponseDateString
	for _, note := range notes {
		res := model.NoteResponseDateString{
			ID:          note.ID,
			Title:       note.Title,
			Description: note.Description,
			Check:       note.Check,
			CreatedAt:   ConvertDatetoString(note.CreatedAt),
			UpdatedAt:   ConvertDatetoString(note.UpdatedAt),
		}
		noteResponses = append(noteResponses, res)
	}

	return noteResponses, nil
}

func (s *NoteServiceImpl) FindNote(ctx context.Context, noteId string) (model.NoteResponseDateString, error) {
	db := s.DB

	notes, err := s.NoteRepository.FindNote(ctx, db, noteId)
	if err != nil {
		return model.NoteResponseDateString{}, err
	}

	res := model.NoteResponseDateString{
		ID:          notes.ID,
		Title:       notes.Title,
		Description: notes.Description,
		Check:       notes.Check,
		CreatedAt:   ConvertDatetoString(notes.CreatedAt),
		UpdatedAt:   ConvertDatetoString(notes.UpdatedAt),
	}

	return res, nil
}

func (s *NoteServiceImpl) CreateNote(ctx context.Context, req model.NoteCreateRequest) (model.NoteResponseDateString, error) {
	db := s.DB

	err := s.Validate.Struct(req)
	exception.PanicIfNeeded(err)

	noteReq := entity.Note{
		Title:       req.Title,
		Description: req.Description,
		Check:       req.Check,
	}

	note, err := s.NoteRepository.CreateNote(ctx, db, noteReq)
	if err != nil {
		return model.NoteResponseDateString{}, err
	}

	res := model.NoteResponseDateString{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		Check:       note.Check,
		CreatedAt:   ConvertDatetoString(note.CreatedAt),
		UpdatedAt:   ConvertDatetoString(note.UpdatedAt),
	}

	return res, nil
}

func (s *NoteServiceImpl) UpdateNote(ctx context.Context, req model.NoteUpdateRequest) (model.NoteResponseDateString, error) {
	db := s.DB
	idNote, err := strconv.Atoi(req.ID)
	if err != nil {
		return model.NoteResponseDateString{}, err
	}
	noteReq := entity.Note{
		ID:          idNote,
		Title:       req.Title,
		Description: req.Description,
		Check:       req.Check,
	}
	res, err := s.NoteRepository.UpdateNote(ctx, db, noteReq)
	if err != nil {
		return model.NoteResponseDateString{}, err
	}
	var response = model.NoteResponseDateString{
		ID:          res.ID,
		Title:       res.Title,
		Description: res.Description,
		Check:       res.Check,
		CreatedAt:   ConvertDatetoString(res.CreatedAt),
		UpdatedAt:   ConvertDatetoString(res.UpdatedAt),
	}
	return response, nil
}

func (s *NoteServiceImpl) DeleteNote(ctx context.Context, id string) error {
	db := s.DB
	err := s.NoteRepository.DeleteNote(ctx, db, id)

	if err != nil {
		return err
	}
	return nil
}

func ConvertDatetoString(t time.Time) string {
	month := t.Format("01")
	var m string
	switch month {
	case "01":
		m = "Januari"
	case "02":
		m = "Februari"
	case "03":
		m = "Maret"
	case "04":
		m = "April"
	case "05":
		m = "Mei"
	case "06":
		m = "Juni"
	case "07":
		m = "Juli"
	case "08":
		m = "Agustus"
	case "09":
		m = "September"
	case "10":
		m = "Oktober"
	case "11":
		m = "November"
	case "12":
		m = "Desember"
	}
	return t.Format("02 " + m + " 2006")
}
