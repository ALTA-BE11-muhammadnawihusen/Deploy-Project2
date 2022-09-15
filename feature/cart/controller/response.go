package controller

import "ecommerce-project/feature/cart/entities"

type CartResponse struct {
	ID           uint
	ProductName  string
	Productimage string
	ProductPrice int
	Quantity     int
	ProductID    uint
}

func CartCoreToResponse(data entities.CoreCart) CartResponse {
	return CartResponse{
		ID:           data.ID,
		ProductName:  data.ProductName,
		Productimage: data.Productimage,
		ProductPrice: data.ProductPrice,
		Quantity:     data.Quantity,
		ProductID:    data.ProductID,
	}
}

func CartCoreToResponseList(data []entities.CoreCart) []CartResponse {
	var list []CartResponse
	for _, v := range data {
		list = append(list, CartCoreToResponse(v))
	}

	return list
}
