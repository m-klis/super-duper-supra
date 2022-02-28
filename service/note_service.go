package service

import (
	"context"
	"noteapp/model"
)

type NoteService interface {
	FindNotes(context.Context) ([]model.NoteResponse, error)
	FindNote(context.Context, string) (model.NoteResponse, error)
}
