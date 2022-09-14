package repository

import (
	"ecommerce-project/feature/product/entities"
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

func (storage *Storage) InsertI(core entities.CoreProduct) (int, error) {
	model := models.CoreToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (storage *Storage) SelectAll(page int) ([]entities.CoreProduct, error) {
	var data []models.Product
	var count = 8 * (page - 1)

	tx := storage.query.Limit(8).Offset(count).Find(&data)
	if tx.Error != nil {
		return []entities.CoreProduct{}, tx.Error
	}
	corelist := models.ProductToCoreList(data)
	return corelist, nil

}

func (storage *Storage) Delete(userid, deleteid int) (string, error) {
	tx := storage.query.Where("user_id = ? and id = ?", userid, deleteid).Delete(&models.Product{})

	if tx.Error != nil || tx.RowsAffected == 0 {
		return "Gagal Menghapus", tx.Error
	}

	return "Sukses Menghapus Barang", nil
}

func (storage *Storage) SelectMyProduct(id, page int) ([]entities.CoreProduct, error) {
	var data []models.Product
	count := 8 * (page - 1)
	tx := storage.query.Limit(8).Offset(count).Find(&data, "user_id = ?", id)
	if tx.Error != nil {
		return []entities.CoreProduct{}, tx.Error
	}
	corelist := models.ProductToCoreList(data)
	return corelist, nil

}

func (storage *Storage) SelectAProduct(idproduct uint) (entities.CoreProduct, error) {
	var data models.Product
	tx := storage.query.Find(&data, "id = ?", idproduct)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return entities.CoreProduct{}, tx.Error
	}

	core := models.ProductToCore(data)
	return core, nil
}

func (storage *Storage) UpdateMyProduct(core entities.CoreProduct, idproduct uint) (string, error) {
	model := models.CoreToModel(core)
	tx := storage.query.Model(&model).Where("id = ?", idproduct).Updates(model)
	if tx.Error != nil || tx.RowsAffected != 1 {
		return "Gagal Update", tx.Error
	}

	return "Sukses Update", nil
}
