package product

import (
	"lami/app/features/users/data"
	"time"
)

type Core struct {
	ID        int
	Name      string
	URL       string
	Price     float32
	Stock     uint
	Detail    string
	StoreName string
	City      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int
	User      data.User
}

type CoreRating struct {
	ID          int
	ProductName string
	URL         string
	ProductID   int
	UserID      int
	Rating      uint
	Review      string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Product     Core
}

type Business interface {
	AddProduct(dataReq Core) (int, error)
	UpdateProduct(dataReq Core, idProduct, idUser int) error
	DeleteProduct(idProduct, idUser int) error
	// SelectProductList() error
	SelectProductbyIDProduct(idProduct int) (Core, error)
	SelectMyProduct(idUser int) ([]Core, error)

	AddProductRating(dataReq CoreRating) (int, error)
	SelectRating(idProduct int) (CoreRating, error)
}

type Data interface {
	AddDataProduct(dataReq Core) (int, error)
	UpdateDataProduct(dataReq Core, idProduct, idUser int) error
	DeleteDataProduct(idProduct, idUser int) error
	// SelectDataProductList() error
	SelectDataProductbyIDProduct(idProduct int) (Core, error)
	SelectDataMyProduct(idUser int) ([]Core, error)

	AddProductDataRating(dataReq CoreRating) (int, error)
	SelectDataRating(idProduct int) (CoreRating, error)
}
