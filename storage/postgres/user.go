package postgres

import (
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) UserByUsername(username string) (*models.User, error) {
	var u models.User
	s.DB.Where(&models.User{Username: username}).First(&u)
	return &u, nil
}

func (s *UserService) AddUser(u *models.User) (*models.User, error) {
	s.DB.Save(&u)
	return u, nil
}
