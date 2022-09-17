package controller

import "ecommerce-project/feature/user/entities"

type Response struct {
	name     string
	email    string
	foto     string
	username string
}

func CoreToResponse(core entities.CoreUser) Response {
	return Response{
		name:     core.Name,
		email:    core.Email,
		foto:     core.Foto,
		username: core.Username,
	}
}
