package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreRecord struct {
	gorm.Model
	ScoreID     uuid.UUID
	UserID      uint
	User        User `gorm:"-"`
	Punctuation uint
	QuestionID  uint
	AnswerID    uint
}

type ScoreRecordService interface {
	SaveScoreRecord(sr *ScoreRecord) (*ScoreRecord, error)
	GetScoreRecords(uuid *uuid.UUID) (*[]ScoreRecord, error)
}

type Score struct {
	gorm.Model
	UserID     uint
	User       User `gorm:"-"`
	ScoreID    uuid.UUID
	TotalScore uint
}

type ScoreService interface {
	SaveScore(s *Score) (*Score, error)
	GetScoreByUUID(uuid *uuid.UUID) (*Score, error)
	GetScoreByID(id uint) (*Score, error)
}
