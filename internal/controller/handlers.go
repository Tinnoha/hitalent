package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testovoe/internal/entity"
	"testovoe/internal/pkg"
	"testovoe/internal/usecase"
	"time"
)

type HTTPHandler struct {
	answer   *usecase.AnswerUseCase
	question *usecase.QuestionUseCase
	logger   pkg.Logger
}

func NewHTTPHandler(answerUC *usecase.AnswerUseCase, questionUC *usecase.QuestionUseCase, logger pkg.Logger) *HTTPHandler {
	return &HTTPHandler{
		answer:   answerUC,
		question: questionUC,
		logger:   logger.WithFields(map[string]interface{}{"layer": "http"}),
	}
}

type ErrorDTO struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (er *ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(er, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(b)
}

func httpError(w http.ResponseWriter, err error, status int) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}

	http.Error(
		w,
		errDTO.ToString(),
		status,
	)
}

// GET All questions input - nothing                  output - json all question
func (h *HTTPHandler) QuestionGetAll(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent())

	start := time.Now()

	questions, err := h.question.GetAll()

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	b, err := json.MarshalIndent(questions, "", "    ")

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	h.logger.Info("questions getet via HTTP",
		"duration", time.Since(start),
	)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error to write answer")
	}
}

// GET 1 question    input - query id                 output - json one question
func (h *HTTPHandler) QuestionGetById(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)

	start := time.Now()

	StringID := r.PathValue("id")
	if StringID == "" {
		httpError(w, errors.New("This id is empty"), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(StringID)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	question, err := h.question.GetByID(id)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(question, "", "    ")

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	h.logger.Info("question geted via HTTP",
		"duration", time.Since(start),
	)

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error to write answer")
	}
}

// POST question     input - json with data question  output - json created question
func (h *HTTPHandler) QuestionCreate(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)
	start := time.Now()

	questionDTO := entity.QuestionDto{}

	err := json.NewDecoder(r.Body).Decode(&questionDTO)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	question, err := h.question.Save(questionDTO)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(question, "", "    ")

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	h.logger.Info("question created via HTTP", "duration", time.Since(start))

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error to write answer")
	}
}

// DELETE question   input - query id                 output - nothing
func (h *HTTPHandler) QuestionDelete(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)
	start := time.Now()

	StringID := r.PathValue("id")
	if StringID != "" {
		id, err := strconv.Atoi(StringID)

		if err != nil {
			httpError(w, err, http.StatusBadRequest)
			return
		}

		err = h.question.Delete(id)

		h.logger.Info("question deleted via HTTP", "duration", time.Since(start))

		if err != nil {
			httpError(w, err, http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	} else {
		httpError(w, errors.New("This id is empty"), http.StatusBadRequest)
	}
}

func (h *HTTPHandler) AnswerGetById(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)
	start := time.Now()

	StringID := r.PathValue("id")
	if StringID == "" {
		httpError(w, errors.New("This id is empty"), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(StringID)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	answer, err := h.answer.GetByID(id)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(answer, "", "    ")

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	h.logger.Info("answer geted via HTTP", "duration", time.Since(start))

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error to write answer")
	}
}

func (h *HTTPHandler) AnswerCreate(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)
	start := time.Now()

	StringID := r.PathValue("id")
	if StringID == "" {
		httpError(w, errors.New("This id is empty"), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(StringID)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	AnswerDTO := entity.AnswerDto{}

	err = json.NewDecoder(r.Body).Decode(&AnswerDTO)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	answer, err := h.answer.Save(AnswerDTO, id)

	if err != nil {
		httpError(w, err, http.StatusBadRequest)
		return
	}

	b, err := json.MarshalIndent(answer, "", "    ")

	if err != nil {
		httpError(w, err, http.StatusInternalServerError)
		return
	}

	h.logger.Info("answer created via HTTP", "duration", time.Since(start))

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Error to write answer")
	}
}

func (h *HTTPHandler) AnswerDelete(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("HTTP request received",
		"method", r.Method,
		"path", r.URL.Path,
		"user_agent", r.UserAgent(),
	)
	start := time.Now()

	StringID := r.PathValue("id")
	if StringID != "" {
		id, err := strconv.Atoi(StringID)

		if err != nil {
			httpError(w, err, http.StatusBadRequest)
			return
		}

		err = h.answer.Delete(id)

		h.logger.Info("answer deleted via HTTP", "duration", time.Since(start))

		if err != nil {
			httpError(w, err, http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	} else {
		httpError(w, errors.New("This id is empty"), http.StatusBadRequest)
	}
}
