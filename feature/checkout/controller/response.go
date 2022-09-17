package controller

import (
	"ecommerce-project/feature/checkout/entities"
	"time"
)

type Response struct {
	qty    int
	total  int
	status string
	time   time.Time
}

func CoreToResponseHist(data entities.CoreOrderHistory) Response {
	return Response{
		qty:    data.Qty,
		total:  data.Total,
		status: data.Status,
		time:   data.Time,
	}
}

func CoreToResponseHistList(data []entities.CoreOrderHistory) []Response {
	var list []Response
	for _, v := range data {
		list = append(list, CoreToResponseHist(v))
	}

	return list
}
