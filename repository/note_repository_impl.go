package repository

import (
	"context"
	"errors"
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

func (repo *noteRepositoryImpl) CreateNote(ctx context.Context, db *sqlx.DB, note entity.Note) (entity.Note, error) {
	stmt, err := db.PrepareNamedContext(ctx, sqlhelper.CreateNote)
	exception.PanicIfNeeded(err)

	var id int
	err = stmt.Get(&id, note)
	exception.PanicIfNeeded(err)

	note.ID = id

	// _, err := db.NamedExecContext(ctx, sqlhelper.CreateNote, note)
	// exception.PanicIfNeeded(err)

	return note, nil
}

func (repo *noteRepositoryImpl) FindNotes(ctx context.Context, db *sqlx.DB) ([]entity.Note, error) {
	var notes []entity.Note

	err := db.SelectContext(ctx, &notes, sqlhelper.FindAll)
	if err != nil {
		return []entity.Note{}, errors.New("failed take data")
	}
	// exception.PanicIfNeeded(err)

	return notes, nil
}

func (repo *noteRepositoryImpl) FindNote(ctx context.Context, db *sqlx.DB, noteId string) (entity.Note, error) {
	var note entity.Note

	err := db.GetContext(ctx, &note, sqlhelper.FindOne+string(noteId))
	if err != nil {
		return entity.Note{}, errors.New("failed take data")
	}
	// exception.PanicIfNeeded(err)

	return note, nil
}
