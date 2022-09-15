package entities

type ServiceInterface interface {
}

type RepositoryInterface interface {
	GetandDelete(userid int) ([]Cart, error)
}
