package models

import "gorm.io/gorm"

type role string

const (
	Admin  role = "admin"
	Player role = "player"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     role `sql:"type:role"`
}

// UserService represent all of methods of the user structure.
type UserService interface {
	UserByUsername(username string) (*User, error)
	AddUser(u *User) (*User, error)
	CheckUserOrEmail(username, email string) (*User, error)
}
