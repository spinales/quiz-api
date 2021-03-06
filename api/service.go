package api

import (
	"github.com/spinales/quiz-api/storage/postgres"
	"gorm.io/gorm"
)

// service represent all of the dependency required for the app
type service struct {
	UserService *postgres.UserService
	// AnswerService   *postgres.AnswerService
	QuestionService *postgres.QuestionService
}

// NewService create new service
func NewService(db *gorm.DB) *service {
	return &service{
		UserService: &postgres.UserService{DB: db},
		// AnswerService:   &postgres.AnswerService{DB: db},
		QuestionService: &postgres.QuestionService{DB: db},
	}
}
