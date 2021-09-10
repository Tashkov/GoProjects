package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
	Date time.Time `json:"date" db:"date"`
	Body string    `json:"body" db:"body"`
}
