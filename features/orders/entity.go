package orders

import (
	_dataCart "lami/app/features/carts/data"
	_dataProduct "lami/app/features/products/data"
	_dataUser "lami/app/features/users/data"
)

type Core struct {
	ID          int
	CartID      []int
	UserID      int
	Receiver    string
	PhoneNumber string
	Address     string
	TotalPrice  uint
	Status      string
	User        _dataUser.User
	Product     _dataProduct.Product
	Cart        _dataCart.Cart
	// Payment     string
}

type CoreDetail struct {
	ID         int
	ProductID  int
	UserID     int
	OrderID    int
	Receiver   string
	Address    string
	Status     string
	TotalPrice uint
	Qty        uint
	User       _dataUser.User
	Product    []Product
}

type Product struct {
	ID   int
	Name string
	Url  string
	Qty  uint
}

type Business interface {
	AddOrder(dataReq Core, idUser int) error
	SelectHistoryOrder(idUser int) (CoreDetail, error)
}

type Data interface {
	AddDataOrder(dataReq Core, idUser int) error
	SelectDataHistoryOrder(idUser int) (CoreDetail, error)
}
