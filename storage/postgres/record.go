package postgres

import (
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type ScoreService struct {
	DB *gorm.DB
}

func (s *ScoreService) SaveScore(sr *models.Score) (*models.Score, error) {
	s.DB.Create(&sr)
	return sr, nil
}

func (s *ScoreService) TopScores() (*[]models.Score, error) {
	var sc []models.Score
	s.DB.Order("total_score asc").Limit(10).Find(&sc)
	for _, v := range sc {
		s.DB.First(&v.User, v.UserID)
	}
	return &sc, nil
}

func (s *ScoreService) GetScoreByUsername(username string) (*[]models.Score, error) {
	var sc []models.Score

	var user models.User
	s.DB.Where(&models.User{Username: username}).First(&user)
	s.DB.Where(&models.Score{UserID: user.ID}).Find(&sc)
	return &sc, nil
}
