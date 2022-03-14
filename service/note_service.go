package service

import (
	"context"
	"noteapp/model"
)

type NoteService interface {
	FindNotes(context.Context) ([]model.NoteResponseDateString, error)
	FindNote(context.Context, string) (model.NoteResponseDateString, error)
	CreateNote(context.Context, model.NoteCreateRequest) (model.NoteResponseDateString, error)
	UpdateNote(context.Context, model.NoteUpdateRequest) (model.NoteResponseDateString, error)
	DeleteNote(context.Context, string) error
}
