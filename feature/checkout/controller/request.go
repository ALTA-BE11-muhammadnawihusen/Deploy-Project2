package controller

import "ecommerce-project/feature/checkout/entities"

type CheckOutRequest struct {
	Address string `json:"address" form:"address"`
	Payment string `json:"payment" form:"payment"`
	Tombol  string `json:"tombol" form:"tombol"`
}

func RequestCheckToCore(data CheckOutRequest) entities.CoreCheckOut {
	return entities.CoreCheckOut{
		Address: data.Address,
		Payment: data.Payment,
		Tombol:  data.Tombol,
	}
}
