package postgres

import (
	"github.com/spinales/quiz-api/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (s *UserService) User(id uint) (*models.User, error) {
	var u models.User
	s.DB.First(&u, id)
	return &u, nil
}

func (s *UserService) UserByUsername(username string) (*models.User, error) {
	var u models.User
	s.DB.Where(&models.User{Username: username}).First(&u)
	return &u, nil
}

func (s *UserService) AddUser(u *models.User) (*models.User, error) {
	s.DB.Create(&u)
	return u, nil
}

func (s *UserService) CheckUserOrEmail(username, email string) bool {
	var u models.User
	s.DB.Where(&models.User{Username: username}).Or(&models.User{Email: email}).First(&u)

	if u.Username != "" || u.Email != "" {
		return true
	}

	return false
}

func (s *UserService) UpdateUser(u *models.User, id uint) (*models.User, error) {
	u.ID = id
	s.DB.Save(&u)
	return u, nil
}

func (s *UserService) UsersByScore() (*[]models.User, error) {
	var users []models.User
	s.DB.Order("score asc").Limit(10).Find(&users)
	return &users, nil
}
