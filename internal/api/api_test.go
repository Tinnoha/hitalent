package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"testovoe/internal/controller"
	"testovoe/internal/entity"
	"testovoe/internal/pkg"
	"testovoe/internal/repositoriy"
	"testovoe/internal/usecase"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestServer(t *testing.T) (*httptest.Server, *gorm.DB) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&repositoriy.Question{}, &repositoriy.Answer{})

	logger := pkg.NewZapLogger()
	questionRepo := repositoriy.NewGormQuestionRepository(db, logger)
	answerRepo := repositoriy.NewGormAnswerRepository(db, logger)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	answerUC := usecase.NewAnswerUseCase(answerRepo, questionRepo)
	handlers := controller.NewHTTPHandler(answerUC, questionUC, logger)
	server := &controller.HTTPServer{Handlers: *handlers}

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := http.NewServeMux()
		router.HandleFunc("GET /question", server.Handlers.QuestionGetAll)
		router.HandleFunc("GET /question/{id}", server.Handlers.QuestionGetById)
		router.HandleFunc("POST /question", server.Handlers.QuestionCreate)
		router.HandleFunc("DELETE /question/{id}", server.Handlers.QuestionDelete)
		router.HandleFunc("GET /answer/{id}", server.Handlers.AnswerGetById)
		router.HandleFunc("POST /question/{id}/answer", server.Handlers.AnswerCreate)
		router.HandleFunc("DELETE /answer/{id}", server.Handlers.AnswerDelete)
		router.ServeHTTP(w, r)
	}))

	return testServer, db
}

func TestQuestionAPI(t *testing.T) {
	server, _ := setupTestServer(t)
	defer server.Close()

	userID := uuid.New()

	t.Run("create question", func(t *testing.T) {
		dto := entity.QuestionDto{
			UserID: userID,
			Text:   "Test question",
		}
		body, _ := json.Marshal(dto)

		resp, err := http.Post(server.URL+"/question", "application/json", bytes.NewReader(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		var question entity.Question
		json.NewDecoder(resp.Body).Decode(&question)
		assert.Equal(t, "Test question", question.Text)
	})

	t.Run("get all questions", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/question")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var questions []entity.Question
		json.NewDecoder(resp.Body).Decode(&questions)
		assert.Greater(t, len(questions), 0)
	})

	t.Run("create answer", func(t *testing.T) {
		answerDTO := entity.AnswerDto{
			UserID: userID,
			Text:   "Test answer",
		}
		body, _ := json.Marshal(answerDTO)

		resp, err := http.Post(server.URL+"/question/1/answer", "application/json", bytes.NewReader(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("get answer", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/answer/1")
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})

	t.Run("delete answer", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", server.URL+"/answer/1", nil)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})

	t.Run("delete question", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", server.URL+"/question/1", nil)
		resp, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestQuestionAPI_Errors(t *testing.T) {
	server, _ := setupTestServer(t)
	defer server.Close()

	userID := uuid.New()

	t.Run("short question text", func(t *testing.T) {
		dto := entity.QuestionDto{
			UserID: userID,
			Text:   "Hi",
		}
		body, _ := json.Marshal(dto)

		resp, err := http.Post(server.URL+"/question", "application/json", bytes.NewReader(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("answer to non-existent question", func(t *testing.T) {
		answerDTO := entity.AnswerDto{
			UserID: userID,
			Text:   "Test answer",
		}
		body, _ := json.Marshal(answerDTO)

		resp, err := http.Post(server.URL+"/question/999/answer", "application/json", bytes.NewReader(body))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
