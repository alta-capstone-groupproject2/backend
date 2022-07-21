package cart

import (
	_dataProduct "lami/app/features/products/data"
	_dataUser "lami/app/features/users/data"
)

type Core struct {
	ID          int
	UserID      int
	ProductID   int
	Qty         int
	URL         string
	ProductName string
	Price       uint
	User        _dataUser.User
	Product     _dataProduct.Product
}

type Business interface {
	AddCart(dataReq Core) (int, error)
	SelectCart(idUser int) ([]Core, error)
	UpdateCart(dataReq Core, idCart int) error
	DeleteCart(idCart, idUser int) error
}

type Data interface {
	AddDataCart(dataReq Core) (int, error)
	SelectDataCart(idUser int) ([]Core, error)
	UpdateDataCart(dataReq Core, idCart int) error
	DeleteDataCart(idCart, idUser int) error
}
