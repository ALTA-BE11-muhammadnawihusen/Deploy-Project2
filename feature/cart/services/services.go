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

func (service *Service) GetMyCart(userId uint) ([]entities.CoreCart, error) {
	list, err := service.do.SelectMyCart(userId)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (service *Service) DeleteFromCart(CartId, UserId uint) (string, error) {
	msg, err := service.do.DeleteFromCart(CartId, UserId)
	return msg, err
}
