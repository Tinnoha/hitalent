package repositoriy

import (
	"testovoe/internal/entity"
	"testovoe/internal/pkg"
	"testovoe/internal/usecase"

	"gorm.io/gorm"
)

type GormAnswerRepository struct {
	db     *gorm.DB
	logger pkg.Logger
}

func NewGormAnswerRepository(db *gorm.DB, logger pkg.Logger) usecase.AnswerRepositoriy {
	return &GormAnswerRepository{
		db:     db,
		logger: logger.WithFields(map[string]interface{}{"component": "answer_repository"}),
	}
}

func (r *GormAnswerRepository) GetByID(id int) (entity.Answer, error) {
	r.logger.Debug("getting answer by ID", "answer_id", id)

	var gormAnswer Answer
	result := r.db.First(&gormAnswer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			r.logger.Warn("answer not found", "answer_id", id)
			return entity.Answer{}, result.Error
		}
		r.logger.Error("failed to get answer", "answer_id", id, "error", result.Error)
		return entity.Answer{}, result.Error
	}

	r.logger.Debug("answer found", "answer_id", id)
	return r.toEntity(gormAnswer), nil
}

func (r *GormAnswerRepository) Save(answer entity.Answer) (entity.Answer, error) {
	r.logger.Debug("saving answer",
		"question_id", answer.QuestionID,
		"user_id", answer.UserID,
		"text_length", len(answer.Text))

	gormAnswer := r.toGormModel(answer)

	result := r.db.Create(&gormAnswer)
	if result.Error != nil {
		r.logger.Error("failed to save answer", "error", result.Error, "question_id", answer.QuestionID)
		return entity.Answer{}, result.Error
	}

	savedAnswer := r.toEntity(gormAnswer)

	r.logger.Info("answer saved successfully",
		"answer_id", gormAnswer.ID,
		"question_id", answer.QuestionID)
	return savedAnswer, nil
}

func (r *GormAnswerRepository) Delete(id int) error {
	r.logger.Debug("deleting answer", "answer_id", id)

	result := r.db.Delete(&Answer{}, id)
	if result.Error != nil {
		r.logger.Error("failed to delete answer", "answer_id", id, "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.logger.Warn("answer not found for deletion", "answer_id", id)
		return gorm.ErrRecordNotFound
	}

	r.logger.Info("answer deleted successfully", "answer_id", id)
	return nil
}

func (r *GormAnswerRepository) toEntity(gormAnswer Answer) entity.Answer {
	return entity.Answer{
		ID:         gormAnswer.ID,
		QuestionID: gormAnswer.QuestionID,
		UserID:     gormAnswer.UserID,
		Text:       gormAnswer.Text,
		CreatedAt:  gormAnswer.CreatedAt,
	}
}

func (r *GormAnswerRepository) toGormModel(answer entity.Answer) Answer {
	return Answer{
		ID:         answer.ID,
		QuestionID: answer.QuestionID,
		UserID:     answer.UserID,
		Text:       answer.Text,
		CreatedAt:  answer.CreatedAt,
	}
}
