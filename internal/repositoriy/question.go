package repositoriy

import (
	"testovoe/internal/entity"
	"testovoe/internal/pkg"
	"testovoe/internal/usecase"

	"gorm.io/gorm"
)

type GormQuestionRepository struct {
	db     *gorm.DB
	logger pkg.Logger
}

func NewGormQuestionRepository(db *gorm.DB, logger pkg.Logger) usecase.QuestionRepositoriy {
	return &GormQuestionRepository{
		db:     db,
		logger: logger.WithFields(map[string]interface{}{"component": "question_repository"}),
	}
}

func (r *GormQuestionRepository) GetAll() ([]entity.Question, error) {
	r.logger.Debug("getting all questions")

	var gormQuestions []Question
	result := r.db.Find(&gormQuestions)
	if result.Error != nil {
		r.logger.Error("failed to get all questions", "error", result.Error)
		return nil, result.Error
	}

	questions := make([]entity.Question, len(gormQuestions))
	for i, gq := range gormQuestions {
		questions[i] = r.toEntity(gq)
	}

	r.logger.Debug("retrieved questions", "count", len(questions))
	return questions, nil
}

func (r *GormQuestionRepository) GetByID(id int) (entity.Question, error) {
	r.logger.Debug("getting question by ID", "question_id", id)

	var gormQuestion Question
	result := r.db.First(&gormQuestion, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			r.logger.Warn("question not found", "question_id", id)
			return entity.Question{}, result.Error
		}
		r.logger.Error("failed to get question", "question_id", id, "error", result.Error)
		return entity.Question{}, result.Error
	}

	r.logger.Debug("question found", "question_id", id)
	return r.toEntity(gormQuestion), nil
}

func (r *GormQuestionRepository) Save(question entity.Question) (entity.Question, error) {
	r.logger.Debug("saving question", "user_id", question.UserID, "text_length", len(question.Text))

	gormQuestion := r.toGormModel(question)

	result := r.db.Create(&gormQuestion)
	if result.Error != nil {
		r.logger.Error("failed to save question", "error", result.Error, "user_id", question.UserID)
		return entity.Question{}, result.Error
	}
	savedQuestion := r.toEntity(gormQuestion)

	r.logger.Info("question saved successfully", "question_id", gormQuestion.ID)
	return savedQuestion, nil
}

func (r *GormQuestionRepository) Delete(id int) error {
	r.logger.Debug("deleting question", "question_id", id)

	result := r.db.Delete(&Question{}, id)
	if result.Error != nil {
		r.logger.Error("failed to delete question", "question_id", id, "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.logger.Warn("question not found for deletion", "question_id", id)
		return gorm.ErrRecordNotFound
	}

	r.logger.Info("question deleted successfully", "question_id", id)
	return nil
}
func (r *GormQuestionRepository) toEntity(gormQuestion Question) entity.Question {
	return entity.Question{
		ID:        gormQuestion.ID,
		UserID:    gormQuestion.UserID,
		Text:      gormQuestion.Text,
		CreatedAt: gormQuestion.CreatedAt,
	}
}

func (r *GormQuestionRepository) toGormModel(question entity.Question) Question {
	return Question{
		ID:        question.ID,
		UserID:    question.UserID,
		Text:      question.Text,
		CreatedAt: question.CreatedAt,
	}
}
