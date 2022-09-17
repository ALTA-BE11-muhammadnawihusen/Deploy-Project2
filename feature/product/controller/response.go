package controller

import "ecommerce-project/feature/product/entities"

type ResponseProduct struct {
	iD       uint
	name     string
	price    int
	foto     string
	quantity uint
}

func CoreToResponse(data entities.CoreProduct) ResponseProduct {
	var response ResponseProduct
	response.iD = data.ID
	response.name = data.Name
	response.price = data.Price
	response.foto = data.Foto
	response.quantity = data.Quantity

	return response
}

func CoreToResponseList(data []entities.CoreProduct) []ResponseProduct {
	var list []ResponseProduct
	for _, v := range data {
		temp := CoreToResponse(v)
		list = append(list, temp)
	}

	return list
}
