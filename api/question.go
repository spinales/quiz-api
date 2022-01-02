package api

import (
	"encoding/json"
	"net/http"

	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) AddQuestion(w http.ResponseWriter, r *http.Request) {
	var q models.Question
	json.NewDecoder(r.Body).Decode(&q)

	_, err := s.service.QuestionService.SaveQuestion(&q)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, q)
}
