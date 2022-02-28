package service

import (
	"context"
	"noteapp/model"
)

type NoteService interface {
	CreateNote(context.Context, model.NoteCreateRequest) (model.NoteResponse, error)
	FindNotes(context.Context) ([]model.NoteResponse, error)
	FindNote(context.Context, string) (model.NoteResponse, error)
}
