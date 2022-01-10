package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Content  string
	ImageUrl string
	Answers  pq.StringArray `gorm:"type:text[]"`
}

type QuestionService interface {
	GetQuestion(id uint) (*Question, error)
	GetQuestions() (*[]Question, error)
	SaveQuestion(q *Question) (*Question, error)
	UpdateQuestion(q *Question, id uint) (*Question, error)
	DeleteQuestion(id uint) error
}
