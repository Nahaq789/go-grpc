package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"size:100;not null"`
	Email string `json:"email" gorm:"size:191;uniqueIndex;not null"`
}

func NewUser(name, email string) *User {
	return &User{Name: name, Email: email}
}
