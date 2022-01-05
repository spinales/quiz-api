package api

import (
	"encoding/json"
	"net/http"

	"github.com/spinales/quiz-api/models"
	"github.com/spinales/quiz-api/util"
)

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	user, err := s.service.UserService.UserByUsername(req.Username)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
		"deletedAt": user.DeletedAt,
		"username":  user.Username,
		"email":     user.Email,
		"role":      user.Role,
	}})
}

func (s *Server) register(w http.ResponseWriter, r *http.Request) {
	var req models.User
	json.NewDecoder(r.Body).Decode(&req)

	if s.service.UserService.CheckUserOrEmail(req.Username, req.Email) {
		util.RespondWithError(w, http.StatusOK, "This username or Email is taken, try with another.")
		return
	}

	_, err := util.ValidateEmail(req.Email)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	hashPassword, err := util.HashPassword(req.Password)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	req.Password = hashPassword

	user, err := s.service.UserService.AddUser(&req)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	token, err := s.tokenMaker.CreateToken(user.Username, s.config.AccessTokenDuration)
	if err != nil {
		util.RespondWithError(w, http.StatusOK, err.Error())
		return
	}

	util.RespondwithJSON(w, http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{
		"createdAt": user.CreatedAt,
		"updatedAt": user.UpdatedAt,
		"deletedAt": user.DeletedAt,
		"username":  user.Username,
		"email":     user.Email,
		"role":      user.Role,
	}})
}
