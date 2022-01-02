package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func (s *Server) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	var req models.Question
	json.NewDecoder(r.Body).Decode(&req)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	q, err := s.service.QuestionService.UpdateQuestion(&req, uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, q)
}

func (s *Server) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	err = s.service.QuestionService.DeleteQuestion(uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, nil)
}
