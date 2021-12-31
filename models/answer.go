package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content  string
	Score    uint
	ImageUrl string
}

type AnswerService interface {
	GetAnswer(id uint) (*Answer, error)
	GetAnswers() (*[]Answer, error)
	SaveAnswer(a *Answer) (*Answer, error)
	UpdateAnswer(a *Answer, id uint) (*Answer, error)
	DeleteAnswer(id uint) error
}
