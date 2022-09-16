package entities

type ServiceInterface interface {
	GetToHistory(userid int, data CoreCheckOut) (string, error)
	GetHistory(userid int) ([]CoreOrderHistory, error)
}

type RepositoryInterface interface {
	GetOutFromCart(userid int, tombol string) ([]Cart, error)
	Insert(data CoreCheckOut, userid int) (string, error)
	SelectHistory(userid int) ([]CoreOrderHistory, error)
}
