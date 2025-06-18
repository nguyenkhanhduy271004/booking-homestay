package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Code        string `gorm:"unique" json:"code"`
	Description string `json:"description"`
	Roles       []Role `gorm:"many2many:role_permissions;"`
}
