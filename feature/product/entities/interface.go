package entities

type ServiceInterface interface {
	AddProductI(CoreProduct) (int, error)
	GetAll(page int) ([]CoreProduct, error)
	GetMyProduct(id, page int) ([]CoreProduct, error)
	UpdateMyProduct(core CoreProduct, idproduct uint) (string, error)
	Delete(userid, deleteid int) (string, error)
	GetAProduct(idproduct int) (CoreProduct, error)
}

type RepositoryInterface interface {
	InsertI(CoreProduct) (int, error)
	SelectAll(page int) ([]CoreProduct, error)
	SelectMyProduct(id, page int) ([]CoreProduct, error)
	UpdateMyProduct(core CoreProduct, idproduct uint) (string, error)
	Delete(userid, deleteid int) (string, error)
	SelectAProduct(idproduct uint) (CoreProduct, error)
}
