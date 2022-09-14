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

	e.POST("/profile/cart/:id", handler.AddCart, middlewares.JWTMiddleware())
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
