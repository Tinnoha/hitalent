package usecase

import (
	"errors"
	"testovoe/internal/entity"
	"time"
)

type AnswerRepositoriy interface {
	GetByID(int) (entity.Answer, error)
	Save(entity.Answer) (entity.Answer, error)
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

func (uc *AnswerUseCase) Save(dto entity.AnswerDto, questionID int) (entity.Answer, error) {
	if 5 > len(dto.Text) {
		return entity.Answer{}, errors.New("Text of Answer is short")
	}

	if len(dto.Text) > 200 {
		return entity.Answer{}, errors.New("Text of Answer is long")
	}

	_, err := uc.questRepo.GetByID(questionID)

	if err != nil {
		return entity.Answer{}, errors.New("This question is not exist")
	}

	answer := entity.Answer{
		QuestionID: questionID,
		UserID:     dto.UserID,
		Text:       dto.Text,
		CreatedAt:  time.Now(),
	}

	return uc.ansRepo.Save(answer)
}

func (uc *AnswerUseCase) GetByID(questionID int) (entity.Answer, error) {
	return uc.ansRepo.GetByID(questionID)
}

func (uc *AnswerUseCase) Delete(answerID int) error {
	return uc.ansRepo.Delete(answerID)
}
