package postgres

import (
	"github.com/google/uuid"
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type ScoreRecordService struct {
	DB *gorm.DB
}

func (s *ScoreRecordService) SaveScoreRecord(sr *models.ScoreRecord) (*models.ScoreRecord, error) {
	s.DB.Create(&sr)
	return sr, nil
}

func (s *ScoreRecordService) GetScoreRecords(uuid *uuid.UUID) (*[]models.ScoreRecord, error) {
	var scr []models.ScoreRecord
	s.DB.Where(&models.ScoreRecord{ScoreID: *uuid}).Find(&scr)
	return &scr, nil
}
