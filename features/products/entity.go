package product

import "lami/app/features/users/data"

type Core struct {
	ID         int
	UserID     int
	Name       string
	URL        string
	Price      float32
	Stock      uint
	Detail     string
	StoreName  string
	City       string
	MeanRating float32
	User       data.User
}

type CoreRating struct {
	ID          int
	UserID      int
	ProductID   int
	ProductName string
	URL         string
	Rating      uint
	Review      string
	User        data.User
	Product     Core
}

type Business interface {
	AddProduct(dataReq Core) (int, error)
	UpdateProduct(dataReq Core, idProduct, idUser int) error
	DeleteProduct(idProduct, idUser int) error
	SelectProductList(limit int, page int, city string, name string) ([]Core, int64, error)
	SelectProductbyIDProduct(idProduct int) (Core, error)
	SelectMyProduct(idUser int) ([]Core, error)

	AddProductRating(dataReq CoreRating) (int, error)
	SelectRating(idProduct int) ([]CoreRating, error)
}

type Data interface {
	AddDataProduct(dataReq Core) (int, error)
	UpdateDataProduct(dataReq Core, idProduct, idUser int) error
	DeleteDataProduct(idProduct, idUser int) error
	SelectProductList(limit int, page int, city string, name string) ([]Core, int64, error)
	SelectDataProductbyIDProduct(idProduct int) (Core, error)
	SelectDataMyProduct(idUser int) ([]Core, error)

	AddProductDataRating(dataReq CoreRating) (int, error)
	SelectDataRating(idProduct int) ([]CoreRating, error)
}
