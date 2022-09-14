package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID         uint
	ProductID      uint
	TotalQty       int
	TotalPrice     int
	AddressRequest uint
	PaymentMethod  string
}
