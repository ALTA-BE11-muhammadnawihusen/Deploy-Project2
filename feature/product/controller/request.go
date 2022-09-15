package controller

import "ecommerce-project/feature/product/entities"

type Request struct {
	Name        string `json:"name" form:"name"`
	Price       int    `json:"price" form:"price"`
	Foto        string `json:"foto" form:"foto"`
	Description string `json:"description" form:"description"`
	CategoryID  uint   `json:"categoryid" form:"categoryid"`
	Quantity    uint   `json:"quantity" form:"quantity"`
}

func (Req *Request) ReqToCore(userid uint) entities.CoreProduct {
	core := entities.CoreProduct{
		Name:        Req.Name,
		Price:       Req.Price,
		Foto:        Req.Foto,
		Description: Req.Description,
		CategoryID:  Req.CategoryID,
		Quantity:    Req.Quantity,
		UserID:      userid,
	}

	return core
}
