package repositoriy

import (
	"time"

	"github.com/google/uuid"
)

type Question struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Text      string    `gorm:"type:text;not null"`
	CreatedAt time.Time
	Answers   []Answer `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE"`
}

type Answer struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	QuestionID int       `gorm:"not null;index"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	Text       string    `gorm:"type:text;not null"`
	CreatedAt  time.Time
	Question   Question `gorm:"foreignKey:QuestionID"`
}
