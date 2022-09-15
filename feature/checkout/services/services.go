package services

import (
	"ecommerce-project/feature/checkout/entities"
)

type Service struct {
	do entities.RepositoryInterface
}

func New(data entities.RepositoryInterface) entities.ServiceInterface {
	return &Service{
		do: data,
	}
}

func (service *Service) GetToHistory(userid int, data entities.CoreCheckOut) (string, error) {
	listCart, err := service.do.GetOutFromCart(userid, data.Tombol)
	if err != nil {
		return "", err
	}
	// fmt.Println("=====Test2=====")
	for _, v := range listCart {
		harga := v.Quantity * v.ProductPrice
		data.Total += harga
		data.Qty += v.Quantity
	}

	msg, err := service.do.Insert(data, userid)
	return msg, err

}

// Qty     int
// Total   int
