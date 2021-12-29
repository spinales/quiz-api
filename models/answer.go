package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content string
	Score   uint
}

type AnswerService interface {
	GetAnswer(id uint) (*Answer, error)
	GetAnswers() (*[]Answer, error)
	SaveAnswer(a *Answer) (*Answer, error)
	DeleteAnswer(id uint) error
}
