package main

import (
	"log"
	"os"
	"testovoe/internal/controller"
	"testovoe/internal/pkg"
	"testovoe/internal/repository"
	"testovoe/internal/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	logger := pkg.NewConsoleLogger()

	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database", "error", err)
		log.Fatal(err)
	}

	// ❌ УБИРАЕМ AutoMigrate - теперь миграции через Goose
	logger.Info("database connected successfully")

	questionRepo := repository.NewGormQuestionRepository(db, logger)
	answerRepo := repository.NewGormAnswerRepository(db, logger)
	questionUC := usecase.NewQuestionUseCase(questionRepo)
	answerUC := usecase.NewAnswerUseCase(answerRepo, questionRepo)
	handlers := controller.NewHTTPHandler(answerUC, questionUC, logger)
	server := &controller.HTTPServer{Handlers: *handlers}

	logger.Info("starting server on :8080")
	if err := server.Run(); err != nil {
		logger.Error("server failed", "error", err)
	}
}

func getDSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "testovoe")

	return "host=" + host + " user=" + user + " password=" + password +
		" dbname=" + dbname + " port=" + port + " sslmode=disable"
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
