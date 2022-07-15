package request

import (
	"lami/app/features/products"
	"lami/app/features/users/data"
	"time"
)

type Product struct {
	Name     string  `json:"name" form:"name"`
	URL      string  `json:"file" form:"file"`
	Price    float32 `json:"price" form:"price"`
	Stock    uint    `json:"stock" form:"stock"`
	Detail   string  `json:"detail" form:"detail"`
	Date     string  `json:"date" form:"date"`
	DateTime time.Time
	UserID   data.User
}

type Rating struct {
	ProductID int    `json:"productID" form:"productID"`
	Rating    uint   `json:"rating" form:"rating"`
	Review    string `json:"review" form:"review"`
	UserID    data.User
}

//	For insert product
func ToCore(productReq Product) product.Core {
	return product.Core{
		URL:    productReq.URL,
		Date:   productReq.DateTime,
		Name:   productReq.Name,
		Price:  productReq.Price,
		Stock:  productReq.Stock,
		Detail: productReq.Detail,
		// UserID: 0,
	}
}

// For update product
func ToCoreUpdate(productReq Product) product.Core {
	return product.Core{
		Date:   productReq.DateTime,
		Price:  productReq.Price,
		Stock:  productReq.Stock,
		Detail: productReq.Detail,
		// UserID: 0,
	}
}

//	For insert rating
func ToCoreRating(ratingReq Rating) product.CoreRating {
	return product.CoreRating{
		ProductID: ratingReq.ProductID,
		Rating:    ratingReq.Rating,
		Review:    ratingReq.Review,
		// UserID:    0,
	}
}
