package repository

import (
	"context"
	"fmt"
	"noteapp/entity"
	"noteapp/exception"

	"github.com/jmoiron/sqlx"
)

type noteRepositoryImpl struct {
}

func NewNoteRepository() NoteRepository {
	return &noteRepositoryImpl{}
}

func (repo *noteRepositoryImpl) FindAll(ctx context.Context, db *sqlx.DB) []entity.Note {
	// named exec
	var notes []entity.Note
	SQL := "select id, title, description, checked, created_at, updated_at from notes"
	// rows, err := db.QueryContext(ctx, SQL)
	err := db.SelectContext(ctx, &notes, SQL)
	exception.PanicIfNeeded(err)
	fmt.Println("rows")

	// defer rows.Close()

	// var notes []entity.Note
	// for rows.Next() {
	// 	note := entity.Note{}
	// 	err := rows.Scan(&note.ID, &note.Title, &note.Description, &note.Check, &note.CreatedAt, &note.UpdatedAt)
	// 	exception.PanicIfNeeded(err)
	// 	notes = append(notes, note)
	// }

	return notes
}