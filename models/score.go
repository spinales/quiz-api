package models

import "github.com/google/uuid"

type ScoreRecord struct {
	ScoreID     uuid.UUID
	UserID      uint
	Punctuation uint
	QuestionID  uint
	AnswerID    uint
}

type ScoreRecordService interface {
	SaveScoreRecord(sr *ScoreRecord) (*ScoreRecord, error)
	GetScoreRecords(uuid *uuid.UUID) (*[]ScoreRecord, error)
}

type Score struct {
	userID     uint
	ScoreID    uuid.UUID
	TotalScore uint
}

type ScoreService interface {
	SaveScore(s *Score) (*Score, error)
	GetScoreByUUID(uuid *uuid.UUID) (*Score, error)
	GetScoreByID(id uint) (*Score, error)
}
