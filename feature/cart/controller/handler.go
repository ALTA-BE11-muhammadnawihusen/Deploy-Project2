package controller

import (
	"ecommerce-project/feature/cart/entities"
	"ecommerce-project/middlewares"
	"ecommerce-project/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Delivery struct {
	FromTo entities.ServiceInterface
}

func New(e *echo.Echo, data entities.ServiceInterface) {
	handler := &Delivery{
		FromTo: data,
	}

	e.POST("/cart/:id", handler.AddCart, middlewares.JWTMiddleware()) //idbarang
	e.GET("/cart", handler.GetMyCart, middlewares.JWTMiddleware())
	e.DELETE("/cart/:id", handler.Delete, middlewares.JWTMiddleware()) //idcart
	e.PUT("/cart/:id", handler.Update, middlewares.JWTMiddleware())
}

func (user *Delivery) AddCart(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	idproduct, err := strconv.Atoi(c.Param("id"))
	// fmt.Println(idproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Parameter must be number"))
	}

	msg, errs := user.FromTo.AddToCart(uint(userid), uint(idproduct))
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusCreated, helper.Success(msg))
}

func (user *Delivery) GetMyCart(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	slicecore, err := user.FromTo.GetMyCart(uint(userid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed("Terjadi Kesalahan"))
	}
	resposeCart := CartCoreToResponseList(slicecore)

	return c.JSON(http.StatusOK, helper.SuccessGet("Sukses Mendapatkan data", resposeCart))
}

func (user *Delivery) Delete(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	cartid, err := strconv.Atoi(c.Param("id"))
	// fmt.Println(idproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Parameter must be number"))
	}

	msg, errs := user.FromTo.DeleteFromCart(uint(cartid), uint(userid))
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}

func (user *Delivery) Update(c echo.Context) error {
	userid := middlewares.ExtractToken(c)
	var add QuantityReq
	errbind := c.Bind(&add)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Tolong Masukkan Jumlah Akhir"))
	}
	cartid, err := strconv.Atoi(c.Param("id"))
	// fmt.Println(idproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.Failed("Parameter must be number"))
	}
	msg, errs := user.FromTo.UpdateTO(cartid, userid, add.Quantity)
	if errs != nil {
		return c.JSON(http.StatusInternalServerError, helper.Failed(msg))
	}

	return c.JSON(http.StatusOK, helper.Success(msg))
}
