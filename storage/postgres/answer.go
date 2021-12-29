package postgres

import (
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type AnswerService struct {
	DB *gorm.DB
}

func (s *AnswerService) GetAnswer(id uint) (*models.Answer, error) {
	var a models.Answer
	s.DB.First(&a, id)
	return &a, nil
}

func (s *AnswerService) GetAnswers() (*[]models.Answer, error) {
	var as []models.Answer // answers
	s.DB.Find(&as)
	return &as, nil
}

func (s *AnswerService) SaveAnswer(a *models.Answer) (*models.Answer, error) {
	s.DB.Save(&a)
	return a, nil
}

func (s *AnswerService) DeleteAnswer(id uint) error {
	s.DB.Delete(&models.Answer{}, id)
	return nil
}
