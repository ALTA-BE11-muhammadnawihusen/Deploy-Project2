package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID       uint
	ProductName  string
	Productimage string
	ProductPrice int
	Quantity     int `gorm:"default:1"`
	ProductID    uint
}
