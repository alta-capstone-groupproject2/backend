package data

import (
	"lami/app/features/products"
	"lami/app/features/users/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID     int       `json:"user_id" form:"user_id"`
	Name       string    `json:"name" form:"name"`
	URL        string    `json:"url" form:"url"`
	Price      uint      `json:"price" form:"price"`
	Stock      uint      `json:"stock" form:"stock"`
	City       string    `json:"city" form:"city"`
	Detail     string    `json:"detail" form:"detail"`
	MeanRating float64   `json:"mean_rating" form:"mean_rating"`
	User       data.User `gorm:"foreignKey:UserID"`
}

type Rating struct {
	gorm.Model
	UserID    int       `json:"user_id" form:"user_id"`
	ProductID int       `json:"productID" form:"productID"`
	Rating    float64   `json:"rating" form:"rating"`
	Review    string    `json:"review" form:"review"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	User      data.User `gorm:"foreignKey:UserID"`
}

func fromCore(core product.Core) Product {
	return Product{
		UserID: core.UserID,
		URL:    core.URL,
		Name:   core.Name,
		Price:  core.Price,
		Stock:  core.Stock,
		Detail: core.Detail,
	}
}

func fromCoreRating(core product.CoreRating) Rating {
	return Rating{
		UserID:    core.UserID,
		ProductID: core.ProductID,
		Rating:    core.Rating,
		Review:    core.Review,
	}
}

func (data *Product) toCorebyProductID() product.Core {
	return product.Core{
		ID:        int(data.ID),
		URL:       data.URL,
		Name:      data.Name,
		StoreName: data.User.StoreName,
		Price:     data.Price,
		City:      data.User.City,
		Detail:    data.Detail,
	}
}

func ToCorebyProductID(data Product) product.Core {
	return data.toCorebyProductID()
}

func (data *Product) toCoreProductList() product.Core {
	return product.Core{
		ID:    int(data.ID),
		Name:  data.Name,
		URL:   data.URL,
		Price: data.Price,
		City:  data.City,
	}
}

func ToCoreListProductList(data []Product) []product.Core {
	res := []product.Core{}
	for key := range data {
		res = append(res, data[key].toCoreProductList())
	}
	return res
}

func (data *Rating) toCoreRating() product.CoreRating {
	return product.CoreRating{
		ID:          int(data.ID),
		URL:         data.Product.URL,
		ProductName: data.Product.Name,
		Rating:      data.Rating,
		Review:      data.Review,
	}
}

func ToCoreRating(data []Rating) []product.CoreRating {
	res := []product.CoreRating{}
	for key := range data {
		res = append(res, data[key].toCoreRating())
	}
	return res
}

// Get MyProduct
func (data *Product) toCoreMyProduct() product.Core {
	return product.Core{
		ID:    int(data.ID),
		URL:   data.URL,
		Name:  data.Name,
		Price: data.Price,
		City:  data.User.City,
	}
}

func ToCoreMyProduct(data []Product) []product.Core {
	res := []product.Core{}
	for key := range data {
		res = append(res, data[key].toCoreMyProduct())
	}
	return res
}
