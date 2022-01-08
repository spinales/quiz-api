package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) AddNewRecord(w http.ResponseWriter, r *http.Request) {
	var records []models.ScoreRecord
	var sum uint
	uuid := uuid.New()
	json.NewDecoder(r.Body).Decode(&records)

	token := r.Context().Value("user").(paseto.JSONToken)
	userID, err := strconv.ParseUint(token.Jti, 10, 64)

	for _, v := range records {
		v.UserID = uint(userID)
		v.ScoreID = uuid
		sum += v.Punctuation
		_, err := s.service.ScoreRecordService.SaveScoreRecord(&v)
		if err != nil {
			util.RespondWithError(w, http.StatusOK, err.Error())
		}
	}

	score, err := s.service.ScoreService.SaveScore(&models.Score{
		UserID:     uint(userID),
		ScoreID:    uuid,
		TotalScore: sum,
	})
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, score)
}

func (s *Server) TopScores(w http.ResponseWriter, r *http.Request) {
	scores, err := s.service.ScoreService.TopScores()
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, scores)
}
