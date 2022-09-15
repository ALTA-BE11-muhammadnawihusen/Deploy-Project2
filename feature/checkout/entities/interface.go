package entities

type ServiceInterface interface {
	GetToHistory(userid int, data CoreCheckOut) (string, error)
}

type RepositoryInterface interface {
	GetOutFromCart(userid int, tombol string) ([]Cart, error)
	Insert(data CoreCheckOut, userid int) (string, error)
}
