package usecase

import (
	"errors"
	"testovoe/internal/entity"
	"time"
)

type QuestionRepositoriy interface {
	GetAll() ([]entity.Question, error)
	GetByID(int) (entity.Question, error)
	Save(entity.Question) (entity.Question, error)
	Delete(int) error
}

type QuestionUseCase struct {
	repo QuestionRepositoriy
}

func NewQuestionUseCase(repo QuestionRepositoriy) *QuestionUseCase {
	return &QuestionUseCase{repo: repo}
}

func (uc *QuestionUseCase) Save(dto entity.QuestionDto) (entity.Question, error) {
	if 5 > len(dto.Text) {
		return entity.Question{}, errors.New("Text of question is short")
	}

	if len(dto.Text) > 200 {
		return entity.Question{}, errors.New("Text of question is long")
	}

	question := entity.Question{
		UserID:    dto.UserID,
		Text:      dto.Text,
		CreatedAt: time.Now(),
	}

	return uc.repo.Save(question)
}

func (uc *QuestionUseCase) GetAll() ([]entity.Question, error) {
	return uc.repo.GetAll()
}

func (uc *QuestionUseCase) GetByID(ID int) (entity.Question, error) {
	return uc.repo.GetByID(ID)
}

func (uc *QuestionUseCase) Delete(ID int) error {
	return uc.repo.Delete(ID)
}
