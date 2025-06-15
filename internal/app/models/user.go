package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
