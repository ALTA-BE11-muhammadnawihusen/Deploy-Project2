package entities

import "ecommerce-project/feature/product/entities"

type ServiceInterface interface {
	AddToCart(userid, productId uint) (string, error)
	GetMyCart(userid uint) ([]CoreCart, error)
	DeleteFromCart(cartid, userid uint) (string, error)
	UpdateTO(cartid, userid, add int) (string, error)
}

type RepositoryInterface interface {
	GetData(productId uint) (entities.CoreProduct, error)
	AddToCart(data CoreCart) (string, error)
	SelectMyCart(userid uint) ([]CoreCart, error)
	DeleteFromCart(cartid, userid uint) (string, error)
	UpdateTO(cartid, userid, add int) (string, error)
}
