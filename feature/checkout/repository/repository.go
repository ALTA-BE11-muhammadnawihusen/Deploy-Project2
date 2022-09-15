package repository

import (
	"ecommerce-project/feature/checkout/entities"
	"ecommerce-project/models"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) entities.RepositoryInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) GetOutFromCart(userid int, tombol string) ([]entities.Cart, error) {
	var model []models.Cart
	// fmt.Println("=====Test3=====")
	tx := storage.query.Find(&model, "user_id = ?", userid)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tombol == "pay" {
		tx = storage.query.Unscoped().Delete(&model, "user_id = ?", userid)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	data := models.ModelCartToCoreList(model)
	list := models.CoreCartToCoreartList(data)

	return list, nil
}

func (storage *Storage) Insert(data entities.CoreCheckOut, userid int) (string, error) {
	// fmt.Println("=====Test4=====")
	model := models.CoreCheckOutToModelCheckOut(data)
	var coh models.OrderHistory
	coh.Qty = model.Qty
	coh.Total = model.Total
	if model.Tombol == "pay" {
		coh.Status = "success"
	} else {
		coh.Status = "canceled"
	}
	model.UserID = uint(userid)
	coh.UserID = uint(userid)

	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "", tx.Error
	}

	tx = storage.query.Create(&coh)
	if tx.Error != nil {
		return "", tx.Error
	}

	return "Sukses ke Database", nil
}

// Success
// Canceled
