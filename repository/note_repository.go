package repository

import (
	"context"
	"noteapp/entity"

	"github.com/jmoiron/sqlx"
)

type NoteRepository interface {
	FindAll(context.Context, *sqlx.DB) []entity.Note
}
