package controller

// import (
// 	"ecommerce-project/feature/checkout/entities"
// 	"ecommerce-project/middlewares"

// 	"github.com/labstack/echo/v4"
// )

// type Delivery struct {
// 	FromTo entities.ServiceInterface
// }

// func New(e *echo.Echo, data entities.ServiceInterface) {
// 	handler := &Delivery{
// 		FromTo: data,
// 	}

// 	e.PUT("/checkout", handler.CheckHist, middlewares.JWTMiddleware())

// }

// func (user *Delivery) CheckHist(c echo.Context) error {
// 	userid := middlewares.ExtractToken(c)
// }
