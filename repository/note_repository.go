package repository

import (
	"context"
	"noteapp/entity"

	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	FindNotes(context.Context, *sqlx.DB) ([]entity.Note, error)
	FindNote(context.Context, *sqlx.DB, string) (entity.Note, error)
}
