package data

import (
	"time"

	"lami/app/features/products"
	"lami/app/features/users/data"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name   string    `json:"name" form:"name"`
	URL    string    `json:"url" form:"url"`
	Price  float32   `json:"price" form:"price"`
	Stock  uint      `json:"stock" form:"stock"`
	Detail string    `json:"detail" form:"detail"`
	UserID int       `json:"user_id" form:"user_id"`
	Date   time.Time `json:"date" form:"date"`
	User   data.User `gorm:"foreignKey:UserID"`
	// Store     string    `json:"store" form:"store"`
	// City      string    `json:"city" form:"city"`
}

type Rating struct {
	gorm.Model
	// URL         string    `json:"url" form:"url"`
	// ProductName string    `json:"name" form:"name"`
	Rating      uint      `json:"rating" form:"rating"`
	Review      string    `json:"review" form:"review"`
	ProductID   int       `json:"productID" form:"productID"`
	UserID      int       `json:"user_id" form:"user_id"`
	Date        time.Time `json:"date" form:"date"`
	Product     Product
	User        data.User
}

func fromCore(core product.Core) Product {
	return Product{
		URL:    core.URL,
		Name:   core.Name,
		Price:  core.Price,
		Stock:  core.Stock,
		Detail: core.Detail,
		Date:   core.Date,
		UserID: core.UserID,
		// CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	}
}

func fromCoreRating(core product.CoreRating) Rating {
	return Rating{
		Rating:    core.Rating,
		Review:    core.Review,
		ProductID: core.ProductID,
		UserID:    core.UserID,
		Date:      core.Date,
		// CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	}
}

func (data *Product) toCore() product.Core {
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

func ToCore(data Product) product.Core {
	return data.toCore()
}

func (data *Rating) toCore() product.CoreRating {
	return product.CoreRating{
		ID:          int(data.ID),
		URL:         data.Product.URL,
		ProductName: data.Product.Name,
		Rating:      data.Rating,
		Review:      data.Review,
	}
}

func ToCoreRating(data Rating) product.CoreRating {
	return data.toCore()
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
