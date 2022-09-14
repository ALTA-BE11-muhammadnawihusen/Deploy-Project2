package services

import (
	"ecommerce-project/feature/cart/entities"
	"ecommerce-project/models"
)

type Service struct {
	do entities.RepositoryInterface
}

func New(data entities.RepositoryInterface) entities.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) AddToCart(userid, productid uint) (string, error) {
	coreProduct, err := service.do.GetData(productid)
	if err != nil {
		return "Terjadi Kesalahan", err
	}
	coreCart := models.CoreProductToCoreCart(coreProduct, userid)
	msg, errs := service.do.AddToCart(coreCart)
	if errs != nil {
		return "Gagal menambahkan ke cart", errs
	}

	return msg, nil
}
