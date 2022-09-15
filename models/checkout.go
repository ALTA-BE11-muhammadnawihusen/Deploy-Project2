package models

import "gorm.io/gorm"

type CheckOut struct {
	gorm.Model
	CartID  Cart
	Address string
	Payment string
	Qty     int
	Total   int
	Tombol  string
}

type OrderHistory struct {
	Foto   string
	Name   string
	Qty    int
	Total  int
	Status string
}
