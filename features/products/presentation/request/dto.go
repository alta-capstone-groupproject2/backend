package request

import (
	"lami/app/features/products"
	"lami/app/features/users/data"
)

type Product struct {
	Name       string  `json:"name" form:"name"`
	URL        string  `json:"file" form:"file"`
	Price      uint    `json:"price" form:"price"`
	Stock      uint    `json:"stock" form:"stock"`
	Detail     string  `json:"detail" form:"detail"`
	MeanRating float64 `json:"mean_rating" form:"mean_rating"`
	UserID     data.User
}

type Rating struct {
	Rating    float64 `json:"rating" form:"rating"`
	Review    string  `json:"review" form:"review"`
	ProductID int     `json:"productID" form:"productID"`
	Product   Product
	UserID    data.User
}

//	For insert product
func ToCore(productReq Product) product.Core {
	return product.Core{
		URL:        productReq.URL,
		Name:       productReq.Name,
		Price:      productReq.Price,
		Stock:      productReq.Stock,
		Detail:     productReq.Detail,
		MeanRating: productReq.MeanRating,
	}
}

// For update product
func ToCoreUpdate(productReq Product) product.Core {
	return product.Core{
		Price:  productReq.Price,
		Stock:  productReq.Stock,
		Detail: productReq.Detail,
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
