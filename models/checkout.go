package models

import (
	"time"

	"gorm.io/gorm"
)

type CheckOut struct {
	gorm.Model
	Address string
	Payment string
	Qty     int
	Total   int
	Tombol  string
	UserID  uint
}

type OrderHistory struct {
	Qty       int
	Total     int
	Status    string
	UserID    uint
	CreatedAt time.Time
}
