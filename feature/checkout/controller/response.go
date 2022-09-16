package controller

import (
	"ecommerce-project/feature/checkout/entities"
	"time"
)

type Response struct {
	Qty    int
	Total  int
	Status string
	Time   time.Time
}

func CoreToResponseHist(data entities.CoreOrderHistory) Response {
	return Response{
		Qty:    data.Qty,
		Total:  data.Total,
		Status: data.Status,
		Time:   data.Time,
	}
}

func CoreToResponseHistList(data []entities.CoreOrderHistory) []Response {
	var list []Response
	for _, v := range data {
		list = append(list, CoreToResponseHist(v))
	}

	return list
}
