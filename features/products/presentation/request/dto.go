package request

import (
	"lami/app/features/products"
	"lami/app/features/users/data"
	"time"
)

type Product struct {
	Name      string  `json:"name" form:"name"`
	URL       string  `json:"file" form:"file"`
	Price     float32 `json:"price" form:"price"`
	Stock     uint    `json:"stock" form:"stock"`
	Detail    string  `json:"detail" form:"detail"`
	Date      string  `json:"date" form:"date"`
	CreateAt  time.Time
	UpdatedAt time.Time
	UserID    data.User
}

type Rating struct {
	ProductID int    `json:"productID" form:"productID"`
	Rating    uint   `json:"rating" form:"rating"`
	Review    string `json:"review" form:"review"`
	Date      string `json:"date" form:"date"`
	CreateAt  time.Time
	UpdatedAt time.Time
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
		CreateAt:  productReq.CreateAt,
		UpdatedAt: productReq.UpdatedAt,
	}
}

// For update product
func ToCoreUpdate(productReq Product) product.Core {
	return product.Core{
		Price:     productReq.Price,
		Stock:     productReq.Stock,
		Detail:    productReq.Detail,
		UpdatedAt: productReq.UpdatedAt,
	}
}

//	For insert rating
func ToCoreRating(ratingReq Rating) product.CoreRating {
	return product.CoreRating{
		ProductID: ratingReq.ProductID,
		Rating:    ratingReq.Rating,
		Review:    ratingReq.Review,
		CreateAt:  ratingReq.CreateAt,
		UpdatedAt: ratingReq.UpdatedAt,
	}
}
