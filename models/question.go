package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content  string
	ImageUrl string
	Answers  []Answer
}

type QuestionService interface {
	GetQuestion(id uint) (*Question, error)
	GetQuestions() (*[]Question, error)
	SaveQuestion(q *Question) (*Question, error)
	UpdateQuestion(q *Question, id uint) (*Question, error)
	DeleteQuestion(id uint) error
}
