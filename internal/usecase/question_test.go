package usecase_test

import (
	"testing"
	"testovoe/internal/entity"
	"testovoe/internal/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuestionRepo struct{ mock.Mock }

func (m *MockQuestionRepo) GetAll() ([]entity.Question, error) {
	args := m.Called()
	return args.Get(0).([]entity.Question), args.Error(1)
}

func (m *MockQuestionRepo) Save(question entity.Question) (entity.Question, error) {
	args := m.Called(question)
	return args.Get(0).(entity.Question), args.Error(1)
}

func (m *MockQuestionRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestQuestionUseCase_Save(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
	uc := usecase.NewQuestionUseCase(mockRepo)
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		dto := entity.QuestionDto{UserID: userID, Text: "Valid question"}
		question := entity.Question{ID: 1, UserID: userID, Text: "Valid question"}
		mockRepo.On("Save", mock.Anything).Return(question, nil)

		result, err := uc.Save(dto)

		assert.NoError(t, err)
		assert.Equal(t, 1, result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("short text", func(t *testing.T) {
		dto := entity.QuestionDto{UserID: userID, Text: "Hi"}
		result, err := uc.Save(dto)
		assert.Error(t, err)
		assert.Equal(t, entity.Question{}, result)
	})
}

func TestQuestionUseCase_GetAll(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
	uc := usecase.NewQuestionUseCase(mockRepo)
	questions := []entity.Question{{ID: 1, Text: "Test"}}
	mockRepo.On("GetAll").Return(questions, nil)

	result, err := uc.GetAll()

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	mockRepo.AssertExpectations(t)
}

func TestQuestionUseCase_Delete(t *testing.T) {
	mockRepo := new(MockQuestionRepo)
	uc := usecase.NewQuestionUseCase(mockRepo)
	mockRepo.On("Delete", 1).Return(nil)

	err := uc.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
