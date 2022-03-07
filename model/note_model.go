package model

import "time"

type NoteCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Check       bool   `json:"checked"`
}

type NoteUpdateRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Check       bool   `json:"checked"`
}

type NoteResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Check       bool      `json:"checked"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
