package postgres

import (
	"github.com/google/uuid"
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

func (s *ScoreService) GetScoreByUUID(uuid *uuid.UUID) (*models.Score, error) {
	var sc models.Score
	s.DB.Where(&models.Score{ScoreID: *uuid}).First(&sc)
	return &sc, nil
}

func (s *ScoreService) GetScoreByID(id uint) (*models.Score, error) {
	var sc models.Score
	s.DB.First(&sc, id)
	return &sc, nil
}
