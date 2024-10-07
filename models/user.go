package models

import (
	"golang.org/x/crypto/bcrypt"
)

//`gorm:"unique;not null"`

type User struct {
	Id       uint   `json:"id"`
	Username string `gorm:"unique,not null" json:"username"`
	Email    string `gorm:"unique,not null" json:"email"`
	Password []byte `json:"-"`
	Token    string `json:"token"`
	IsAdmin  uint   `gorm:"default:0" json:"is_admin"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
