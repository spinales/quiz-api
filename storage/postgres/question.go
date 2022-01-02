package postgres

import (
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type QuestionService struct {
	DB *gorm.DB
}

func (s *QuestionService) GetQuestion(id uint) (*models.Question, error) {
	var q models.Question
	s.DB.First(&q, id)
	s.DB.Where(&models.Answer{QuestionID: q.ID}).Find(&q.Answers)
	return &q, nil
}

func (s *QuestionService) GetQuestions() (*[]models.Question, error) {
	var qs []models.Question
	s.DB.Find(&qs)
	return &qs, nil
}

func (s *QuestionService) SaveQuestion(q *models.Question) (*models.Question, error) {
	s.DB.Create(&q)
	return q, nil
}

func (s *QuestionService) UpdateQuestion(q *models.Question, id uint) (*models.Question, error) {
	q.ID = id
	s.DB.Save(&q)
	return q, nil
}

func (s *QuestionService) DeleteQuestion(id uint) error {
	s.DB.Delete(&models.Question{}, id)
	return nil
}
