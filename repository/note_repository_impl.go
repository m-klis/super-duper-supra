package repository

import (
	"context"
	"noteapp/entity"
	"noteapp/exception"
	"noteapp/repository/sqlhelper"

	"github.com/jmoiron/sqlx"
)

type noteRepositoryImpl struct {
}

func NewNoteRepository() NoteRepository {
	return &noteRepositoryImpl{}
}

func (repo *noteRepositoryImpl) FindNotes(ctx context.Context, db *sqlx.DB) ([]entity.Note, error) {
	var notes []entity.Note
	err := db.SelectContext(ctx, &notes, sqlhelper.FindAll)
	exception.PanicIfNeeded(err)
	return notes, nil
}

func (repo *noteRepositoryImpl) FindNote(ctx context.Context, db *sqlx.DB, noteId string) (entity.Note, error) {
	var note entity.Note
	err := db.GetContext(ctx, &note, sqlhelper.FindOne+string(noteId))
	if err != nil {
		return entity.Note{}, err
	}
	return note, nil
}
