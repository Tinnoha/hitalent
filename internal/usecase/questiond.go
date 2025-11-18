package usecase

import (
	"errors"
	"testovoe/internal/entity"
	"time"
)

type QuestionRepositoriy interface {
	GetAll() ([]entity.Question, error)
	GetByID(int) (entity.Question, error)
	Save(entity.Question) error
	Delete(int) error
}

type QuestionUseCase struct {
	repo QuestionRepositoriy
}

func NewQuestionUseCase(repo QuestionRepositoriy) *QuestionUseCase {
	return &QuestionUseCase{repo: repo}
}

func (uc *QuestionUseCase) Save(dto entity.QuestionDto) error {
	if 5 > len(dto.Text) {
		return errors.New("Text of question is short")
	}

	if len(dto.Text) > 200 {
		return errors.New("Text of question is long")
	}

	question := entity.Question{
		UserID:    dto.UserID,
		Text:      dto.Text,
		CreatedAt: time.Now(),
	}

	return uc.repo.Save(question)
}
