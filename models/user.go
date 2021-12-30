package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Role     string
}

// UserService represent all of methods of the user structure.
type UserService interface {
	UserByUsername(username string) (*User, error)
	AddUser(u *User) (*User, error)
}
