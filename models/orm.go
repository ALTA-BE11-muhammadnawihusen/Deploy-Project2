package models

import (
	auth "ecommerce-project/feature/auth/entities"
	cart "ecommerce-project/feature/cart/entities"
	check "ecommerce-project/feature/checkout/entities"
	product "ecommerce-project/feature/product/entities"
	"ecommerce-project/feature/user/entities"
)

func CoreToModel(core product.CoreProduct) Product {
	return Product{
		Name:        core.Name,
		Price:       core.Price,
		Description: core.Description,
		Foto:        core.Foto,
		UserID:      core.UserID,
		CategoryID:  core.CategoryID,
		Quantity:    core.Quantity,
	}
}

func RegCoreToModel(core auth.CoreRegister) User {
	return User{
		Username: core.Username,
		Email:    core.Email,
		Password: core.Password,
	}
}

func ModelToCore(model User) entities.CoreUser {
	return entities.CoreUser{
		Name:     model.Name,
		Email:    model.Email,
		Foto:     model.Foto,
		Username: model.Username,
	}
}

func ProductToCore(data Product) product.CoreProduct {
	return product.CoreProduct{
		ID:         data.ID,
		Name:       data.Name,
		Price:      data.Price,
		Foto:       data.Foto,
		UserID:     data.UserID,
		CategoryID: data.CategoryID,
		Quantity:   data.Quantity,
	}
}

func ProductToCoreList(data []Product) []product.CoreProduct {
	var corelist []product.CoreProduct
	for _, v := range data {
		temp := ProductToCore(v)
		corelist = append(corelist, temp)
	}

	return corelist
}

func CoreUserToModel(data entities.CoreUser) User {
	return User{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Foto:     data.Foto,
	}
}

func CoreProductToCoreCart(data product.CoreProduct, userId uint) cart.CoreCart {
	return cart.CoreCart{
		UserID:       userId,
		ProductName:  data.Name,
		Productimage: data.Foto,
		ProductPrice: data.Price,
		ProductID:    data.ID,
	}
}

func CoreCartToModel(data cart.CoreCart) Cart {
	return Cart{
		UserID:       data.UserID,
		ProductName:  data.ProductName,
		Productimage: data.Productimage,
		ProductPrice: data.ProductPrice,
		ProductID:    data.ProductID,
	}
}

func ModelCartToCore(data Cart) cart.CoreCart {
	return cart.CoreCart{
		ID:           data.ID,
		ProductName:  data.ProductName,
		Productimage: data.Productimage,
		ProductPrice: data.ProductPrice,
		Quantity:     data.Quantity,
		ProductID:    data.ProductID,
	}
}

func ModelCartToCoreList(data []Cart) []cart.CoreCart {
	var list []cart.CoreCart
	for _, v := range data {
		list = append(list, ModelCartToCore(v))
	}

	return list
}

func CoreCartToCoreCart(data cart.CoreCart) check.Cart {
	return check.Cart{
		UserID:       data.UserID,
		ProductName:  data.ProductName,
		Productimage: data.Productimage,
		ProductPrice: data.ProductPrice,
		Quantity:     data.Quantity,
		ProductID:    data.ProductID,
	}
}
func CoreCartToCoreartList(data []cart.CoreCart) []check.Cart {
	var list []check.Cart
	for _, v := range data {
		list = append(list, CoreCartToCoreCart(v))
	}
	return list
}

func CoreCheckOutToModelCheckOut(data check.CoreCheckOut) CheckOut {
	return CheckOut{
		Address: data.Address,
		Payment: data.Payment,
		Qty:     data.Qty,
		Total:   data.Total,
		Tombol:  data.Tombol,
	}
}

func ModelsHistToCore(data OrderHistory) check.CoreOrderHistory {
	return check.CoreOrderHistory{
		Qty:    data.Qty,
		Total:  data.Total,
		Status: data.Status,
		Time:   data.CreatedAt,
	}
}

func ModelsHistToCoreList(data []OrderHistory) []check.CoreOrderHistory {
	var hist []check.CoreOrderHistory
	for _, v := range data {
		hist = append(hist, ModelsHistToCore(v))
	}

	return hist
}
