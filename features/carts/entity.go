package cart

import "lami/app/features/products/data"

type Core struct {
	ID        int
	UserID      int
	ProductID   int
	Qty         int
	URL         string
	ProductName string
	Price       float32
	Product     data.Product
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
