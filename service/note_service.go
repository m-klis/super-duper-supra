package service

import (
	"context"
	"noteapp/model"
)

type NoteService interface {
	FindNotes(context.Context) ([]model.NoteResponse, error)
	FindNote(context.Context, string) (model.NoteResponse, error)
	CreateNote(context.Context, model.NoteCreateRequest) (model.NoteResponse, error)
	UpdateNote(context.Context, model.NoteUpdateRequest) (model.NoteResponse, error)
	DeleteNote(context.Context, string) error
}
