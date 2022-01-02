package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) AddAnswer(w http.ResponseWriter, r *http.Request) {
	var a models.Answer
	json.NewDecoder(r.Body).Decode(&a)

	_, err := s.service.AnswerService.SaveAnswer(&a)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, a)
}

func (s *Server) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	var req models.Answer
	json.NewDecoder(r.Body).Decode(&req)

	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	a, err := s.service.AnswerService.UpdateAnswer(&req, uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, a)
}

func (s *Server) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	err = s.service.AnswerService.DeleteAnswer(uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, nil)
}

func (s *Server) GetAnswer(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	a, err := s.service.AnswerService.GetAnswer(uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, a)
}

func (s *Server) GetAnswers(w http.ResponseWriter, r *http.Request) {
	answers, err := s.service.AnswerService.GetAnswers()
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, answers)
}
