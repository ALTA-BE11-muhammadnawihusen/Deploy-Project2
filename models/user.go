package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Foto     string
	Username string `gorm:"unique"`
	Products []Product
}

// `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
