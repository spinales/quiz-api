package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

// AddQuestion endpoint para agregar una pregunta nueva, el front envia
// un json con la estructura de la preguta, la pregunta en si, la url de
// una imagen y un array de strings que son las respuestas.
func (s *Server) AddQuestion(w http.ResponseWriter, r *http.Request) {
	var q models.Question
	json.NewDecoder(r.Body).Decode(&q)

	_, err := s.service.QuestionService.SaveQuestion(&q)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, q)
}

// UpdateQuestion endpoint para actualizar una pregunta, similar a
// AddQuestion pero en la url se agrega el id de la pregunta aactualizar.
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

// DeleteQuestion endpoint para borrar una pregunta, se recibe el id
// y se elimina el la pregunta que corresponda al id mandado.
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

// GetQuestion endpoint para retornar la pregunta en base al id recibido
func (s *Server) GetQuestion(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, "Cannot parse this param, "+err.Error())
	}

	q, err := s.service.QuestionService.GetQuestion(uint(id))
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, q)
}

// GetQuestions endpoint para retornar todas las preguntas existentes
func (s *Server) GetQuestions(w http.ResponseWriter, r *http.Request) {
	qs, err := s.service.QuestionService.GetQuestions()
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
	}

	util.RespondwithJSON(w, http.StatusOK, qs)
}
