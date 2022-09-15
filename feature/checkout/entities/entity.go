package entities

type CoreCheckOut struct {
	ID      uint
	Address string
	Payment string
	Qty     int
	Total   int
	Tombol  string
	UserID  uint
}

type CoreOrderHistory struct {
	Qty    int
	Total  int
	Status string
	UserID uint
}

type Cart struct {
	UserID       uint
	ProductName  string
	Productimage string
	ProductPrice int
	Quantity     int
	ProductID    uint
}
