package entity

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID        int       `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type QuestionDto struct {
	UserID uuid.UUID `json:"user_id"`
	Text   string    `json:"text"`
}
