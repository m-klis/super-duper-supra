package entity

import (
	"time"
)

type Note struct {
	ID          int
	Title       string
	Description string
	Check       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
