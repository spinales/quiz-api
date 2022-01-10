package api

import (
	"encoding/json"
	"net/http"

	"github.com/o1egl/paseto"
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) AddNewRecord(w http.ResponseWriter, r *http.Request) {
	var record models.User
	json.NewDecoder(r.Body).Decode(&record)

	token := r.Context().Value("pasetoStruct").(*paseto.JSONToken)
	user, err := s.service.UserService.UserByUsername(token.Issuer)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}
	user.Score = record.Score

	_, err = s.service.UserService.UpdateUser(user, user.ID)
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
