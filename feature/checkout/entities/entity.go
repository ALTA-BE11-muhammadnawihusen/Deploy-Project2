package entities

type CheckOut struct {
	ID      uint
	CartID  Cart
	Address string
	Payment string
	Qty     int
	Total   int
	Tombol  string
}

type OrderHistory struct {
	Foto   string
	Name   string
	Qty    int
	Total  int
	Status string
}

type Cart struct {
	UserID       uint
	ProductName  string
	Productimage string
	ProductPrice int
	Quantity     int
	ProductID    uint
}
