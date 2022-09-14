package entities

import "ecommerce-project/feature/product/entities"

type ServiceInterface interface {
	AddToCart(userid, productId uint) (string, error)
}

type RepositoryInterface interface {
	GetData(productId uint) (entities.CoreProduct, error)
	AddToCart(data CoreCart) (string, error)
}
