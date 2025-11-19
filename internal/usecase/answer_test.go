package usecase_test

import (
	"errors"
	"testing"
	"testovoe/internal/entity"
	"testovoe/internal/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAnswerRepo struct{ mock.Mock }

func (m *MockAnswerRepo) GetByID(id int) (entity.Answer, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Answer), args.Error(1)
}

func (m *MockAnswerRepo) Save(answer entity.Answer) (entity.Answer, error) {
	args := m.Called(answer)
	return args.Get(0).(entity.Answer), args.Error(1)
}

func (m *MockAnswerRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockQuestionRepo) GetByID(id int) (entity.Question, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Question), args.Error(1)
}

func TestAnswerUseCase_Save(t *testing.T) {
	mockAnswerRepo := new(MockAnswerRepo)
	mockQuestionRepo := new(MockQuestionRepo)
	uc := usecase.NewAnswerUseCase(mockAnswerRepo, mockQuestionRepo)
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		question := entity.Question{ID: 1}
		answer := entity.Answer{ID: 1, QuestionID: 1, UserID: userID, Text: "Valid answer"}

		mockQuestionRepo.On("GetByID", 1).Return(question, nil)
		mockAnswerRepo.On("Save", mock.Anything).Return(answer, nil)

		dto := entity.AnswerDto{UserID: userID, Text: "Valid answer"}
		result, err := uc.Save(dto, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, result.ID)
		mockQuestionRepo.AssertExpectations(t)
		mockAnswerRepo.AssertExpectations(t)
	})

	t.Run("question not found", func(t *testing.T) {
		mockQuestionRepo.On("GetByID", 999).Return(entity.Question{}, errors.New("not found"))

		dto := entity.AnswerDto{UserID: userID, Text: "Valid answer"}
		result, err := uc.Save(dto, 999)

		assert.Error(t, err)
		assert.Equal(t, entity.Answer{}, result)
	})
}
