package repository

import (
	"ecommerce-project/feature/cart/entities"
	product "ecommerce-project/feature/product/entities"
	"ecommerce-project/models"
	"fmt"

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

func (storage *Storage) GetData(productid uint) (product.CoreProduct, error) {
	var model models.Product
	tx := storage.query.Find(&model, "id = ?", productid)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return product.CoreProduct{}, tx.Error
	}
	core := models.ProductToCore(model)
	fmt.Println(core)
	return core, nil
}

func (storage *Storage) AddToCart(data entities.CoreCart) (string, error) {
	model := models.CoreCartToModel(data)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return "Gagal Menambahkan ke Cart", tx.Error
	}

	fmt.Println(model)
	return "Sukses Menambahkan ke Cart", nil
}

func (storage *Storage) SelectMyCart(userId uint) ([]entities.CoreCart, error) {
	var data []models.Cart
	tx := storage.query.Find(&data, "user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	corelist := models.ModelCartToCoreList(data)
	return corelist, nil
}

func (storage *Storage) DeleteFromCart(cartId, userId uint) (string, error) {
	tx := storage.query.Where("user_id = ? and id = ?", userId, cartId).Delete(&models.Cart{})
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Menghapus", tx.Error
	}

	return "Sukses Menghapus Dari Cart", nil
}

func (storage *Storage) UpdateTO(cartid, userid, add int) (string, error) {
	// var qty models.Cart
	tx := storage.query.Model(&models.Cart{}).Where("id = ? and user_id = ?", cartid, userid).Update("quantity", add)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Terjadi Kesalahan", tx.Error
	}

	return "Sukses Update Data", nil
}

// model := models.CoreToModel(core)
// tx := storage.query.Model(&model).Where("id = ?", idproduct).Updates(model)
// if tx.Error != nil || tx.RowsAffected != 1 {
// 	return "Gagal Update", tx.Error
// }
