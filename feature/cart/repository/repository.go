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

// GetData(productId uint) (entities.CoreProduct, error)
// AddToCart(data CoreCart) (string, error)
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
