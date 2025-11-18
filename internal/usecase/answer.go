package usecase

import (
	"errors"
	"testovoe/internal/entity"
	"time"
)

type AnswerRepositoriy interface {
	GetAll() ([]entity.Answer, error)
	GetByID(int) (entity.Answer, error)
	Save(entity.Answer) error
	Delete(int) error
}

type AnswerUseCase struct {
	ansRepo   AnswerRepositoriy
	questRepo QuestionRepositoriy
}

func NewAnswerUseCase(ansrepo AnswerRepositoriy, quest QuestionRepositoriy) *AnswerUseCase {
	return &AnswerUseCase{
		ansRepo:   ansrepo,
		questRepo: quest,
	}
}

func (uc *AnswerUseCase) Save(dto entity.AnswerDto) error {
	if 5 > len(dto.Text) {
		return errors.New("Text of Answer is short")
	}

	if len(dto.Text) > 200 {
		return errors.New("Text of Answer is long")
	}

	_, err := uc.questRepo.GetByID(dto.QuestionID)

	if err != nil {
		return errors.New("This question is not exist")
	}

	answer := entity.Answer{
		QuestionID: dto.QuestionID,
		UserID:     dto.UserID,
		Text:       dto.Text,
		CreatedAt:  time.Now(),
	}

	return uc.ansRepo.Save(answer)
}
