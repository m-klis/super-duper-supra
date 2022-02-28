package service

import (
	"context"
	"noteapp/model"
)

type NoteService interface {
	FindAll(context.Context) []model.NoteResponse
}
