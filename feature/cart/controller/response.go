package controller

import "ecommerce-project/feature/cart/entities"

type CartResponse struct {
	iD           uint
	productName  string
	productimage string
	productPrice int
	quantity     int
	productID    uint
}

func CartCoreToResponse(data entities.CoreCart) CartResponse {
	return CartResponse{
		iD:           data.ID,
		productName:  data.ProductName,
		productimage: data.Productimage,
		productPrice: data.ProductPrice,
		quantity:     data.Quantity,
		productID:    data.ProductID,
	}
}

func CartCoreToResponseList(data []entities.CoreCart) []CartResponse {
	var list []CartResponse
	for _, v := range data {
		list = append(list, CartCoreToResponse(v))
	}

	return list
}
