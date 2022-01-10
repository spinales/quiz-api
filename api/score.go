package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/o1egl/paseto"
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) AddNewRecord(w http.ResponseWriter, r *http.Request) {
	var record models.User
	json.NewDecoder(r.Body).Decode(&record)

	token := r.Context().Value("user").(paseto.JSONToken)
	userID, err := strconv.ParseUint(token.Jti, 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	user, err := s.service.UserService.User(uint(userID))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}
	user.Score = record.Score

	_, err = s.service.UserService.UpdateUser(user, uint(userID))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, user)
}

func (s *Server) TopScores(w http.ResponseWriter, r *http.Request) {
	scores, err := s.service.UserService.UsersByScore()
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, scores)
}
