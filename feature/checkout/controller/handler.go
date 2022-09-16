package controller

import (
	"ecommerce-project/feature/checkout/entities"
	"ecommerce-project/middlewares"
	"ecommerce-project/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	FromTo entities.ServiceInterface
}

func New(e *echo.Echo, data entities.ServiceInterface) {
	handler := &Delivery{
		FromTo: data,
	}

	e.POST("/checkout", handler.InsertCheckHist, middlewares.JWTMiddleware())
	e.GET("/checkout", handler.CheckHist)

}

func (user *Delivery) InsertCheckHist(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var check CheckOutRequest
	errb := c.Bind(&check)
	if errb != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Gagal Bind Data"))
	}

	if check.Tombol == "pay" || check.Tombol == "cancel" {
		// fmt.Println("=====Test1=====")
		corecheck := RequestCheckToCore(check)
		msg, err := user.FromTo.GetToHistory(userid, corecheck)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, helper.Success(msg))
	}

	return c.JSON(http.StatusBadRequest, helper.Failed("Tolong pilih pay Or cancel"))
}

func (user *Delivery) CheckHist(c echo.Context) error {
	userid := middlewares.ExtractToken(c)

	list, er := user.FromTo.GetHistory(userid)
	if er != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
	}

	coreres := CoreToResponseHistList(list)
	return c.JSON(http.StatusOK, helper.SuccessGet("behasil mendapatkan data", coreres))
}
