package request

import (
	"lami/app/features/products"
	"lami/app/features/users/data"
)

type Product struct {
	Name   string  `json:"name" form:"name"`
	URL    string  `json:"file" form:"file"`
	Price  float32 `json:"price" form:"price"`
	Stock  uint    `json:"stock" form:"stock"`
	Detail string  `json:"detail" form:"detail"`
	UserID data.User
}

type Rating struct {
	Rating    uint   `json:"rating" form:"rating"`
	Review    string `json:"review" form:"review"`
	ProductID int    `json:"productID" form:"productID"`
	Product   Product
	UserID    data.User
}

//	For insert product
func ToCore(productReq Product) product.Core {
	return product.Core{
		URL:       productReq.URL,
		Name:      productReq.Name,
		Price:     productReq.Price,
		Stock:     productReq.Stock,
		Detail:    productReq.Detail,
	}
}

// For update product
func ToCoreUpdate(productReq Product) product.Core {
	return product.Core{
		Price:     productReq.Price,
		Stock:     productReq.Stock,
		Detail:    productReq.Detail,
	}
}

//	For insert rating
func ToCoreRating(ratingReq Rating) product.CoreRating {
	return product.CoreRating{
		ProductID: ratingReq.ProductID,
		Rating:    ratingReq.Rating,
		Review:    ratingReq.Review,
	}
}
