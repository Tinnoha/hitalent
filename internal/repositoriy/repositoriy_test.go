// internal/repository/question_repo_test.go
package repositoriy_test

import (
	"testing"
	"testovoe/internal/entity"
	"testovoe/internal/repositoriy"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&repositoriy.Question{})
	return db
}

func TestQuestionRepository_Save(t *testing.T) {
	db := setupTestDB(t)
	repo := repositoriy.NewGormQuestionRepository(db, nil)
	userID := uuid.New()

	question := entity.Question{
		UserID: userID,
		Text:   "Test question",
	}

	saved, err := repo.Save(question)
	assert.NoError(t, err)
	assert.Equal(t, "Test question", saved.Text)
	assert.NotZero(t, saved.ID)
}

func TestQuestionRepository_GetByID(t *testing.T) {
	db := setupTestDB(t)
	repo := repositoriy.NewGormQuestionRepository(db, nil)
	userID := uuid.New()

	question := entity.Question{UserID: userID, Text: "Test"}
	saved, _ := repo.Save(question)

	found, err := repo.GetByID(saved.ID)
	assert.NoError(t, err)
	assert.Equal(t, saved.ID, found.ID)
}

func TestQuestionRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := repositoriy.NewGormQuestionRepository(db, nil)
	userID := uuid.New()

	repo.Save(entity.Question{UserID: userID, Text: "Q1"})
	repo.Save(entity.Question{UserID: userID, Text: "Q2"})

	questions, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, questions, 2)
}
