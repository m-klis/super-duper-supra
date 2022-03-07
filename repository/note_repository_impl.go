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

func (repo *noteRepositoryImpl) FindNotes(ctx context.Context, db *sqlx.DB) ([]entity.Note, error) {
	var notes []entity.Note

	err := db.SelectContext(ctx, &notes, sqlhelper.FindAll)
	if err != nil {
		return []entity.Note{}, errors.New("failed take data")
	}

	return notes, nil
}

func (repo *noteRepositoryImpl) FindNote(ctx context.Context, db *sqlx.DB, noteId string) (entity.Note, error) {
	var note entity.Note

	err := db.GetContext(ctx, &note, sqlhelper.FindOne+string(noteId))
	if err != nil {
		return entity.Note{}, errors.New("failed take data")
	}

	return note, nil
}

func (repo *noteRepositoryImpl) CreateNote(ctx context.Context, db *sqlx.DB, note entity.Note) (entity.Note, error) {
	stmt, err := db.PrepareNamedContext(ctx, sqlhelper.CreateNote)
	exception.CheckError(err)
	if err != nil {
		return entity.Note{}, err
	}

	var id = entity.Note{}

	err = stmt.Get(&id, note)
	if err != nil {
		return entity.Note{}, err
	}

	note.ID = id.ID
	note.CreatedAt = id.CreatedAt
	note.UpdatedAt = id.UpdatedAt

	return note, nil
}

func (repo *noteRepositoryImpl) UpdateNote(ctx context.Context, db *sqlx.DB, note entity.Note) (entity.Note, error) {
	sqlRows, err := db.QueryContext(ctx, sqlhelper.UpdateNote,
		note.Title, note.Description, note.Check, note.ID)
	if err != nil {
		return entity.Note{}, err
	}
	defer sqlRows.Close()

	for sqlRows.Next() {
		err = sqlRows.Scan(&note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			return entity.Note{}, err
		}
	}

	return note, nil
}

func (repo *noteRepositoryImpl) DeleteNote(ctx context.Context, db *sqlx.DB, id string) error {
	// db.ExecContext(ctx)
	return nil
}
