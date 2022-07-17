package orders

import (
	_dataCart "lami/app/features/carts/data"
	_dataProduct "lami/app/features/products/data"
)

type Core struct {
	ID          int
	UserID      int
	Receiver    string
	PhoneNumber string
	Address     string
	Price       float32
	Status      string
	Product     _dataProduct.Product
	Cart        _dataCart.Cart
	// Payment     string
}

type CoreDetail struct {
	ID          int
	OrderID     int
	ProductName string
	Price       float32
	Qty         uint
}

type CoreCart struct {
	Cart _dataCart.Cart
}

type Business interface {
	AddOrder(dataReq Core, idUser int) (int, error)
}

type Data interface {
	AddDataOrder(dataReq Core, idUser int) (int, error)
}
