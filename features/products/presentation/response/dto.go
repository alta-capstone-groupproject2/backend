package response

import (
	"lami/app/features/products"
)

type Product struct {
	ID         int     `json:"product_id" form:"product_id"`
	URL        string  `json:"url" form:"url"`
	Name       string  `json:"name" form:"name"`
	StoreName  string  `json:"store_name" form:"store_name"`
	Price      uint    `json:"price" form:"price"`
	City       string  `json:"city" form:"city"`
	MeanRating float64 `json:"mean_rating" form:"mean_rating"`
	Detail     string  `json:"detail" form:"detail"`
}

type MyProduct struct {
	ID    int    `json:"product_id" form:"product_id"`
	URL   string `json:"url" form:"url"`
	Name  string `json:"name" form:"name"`
	Price uint   `json:"price" form:"price"`
	City  string `json:"city" form:"city"`
}

type ProductRating struct {
	ID          int     `json:"rating_id" form:"rating_id"`
	URL         string  `json:"url" form:"url"`
	ProductName string  `json:"name" form:"name"`
	Rating      float64 `json:"rating" form:"rating"`
	Review      string  `json:"review" form:"review"`
}

type ProductList struct {
	ID    int    `json:"product_id" form:"product_id"`
	URL   string `json:"url" form:"url"`
	Name  string `json:"name" form:"name"`
	Price uint   `json:"price" form:"price"`
	City  string `json:"city" form:"city"`
}

//	GetProductList
func FromCoreProductList(data product.Core) ProductList {
	return ProductList{
		ID:    data.ID,
		URL:   data.URL,
		Name:  data.Name,
		Price: data.Price,
		City:  data.City,
	}
}

func FromCoreListProductList(data []product.Core) []ProductList {
	res := []ProductList{}
	for key := range data {
		res = append(res, FromCoreProductList(data[key]))
	}
	return res
}

//	GetProductbyIDProduct
func FromCorebyIDProduct(core product.Core) Product {
	return Product{
		ID:         core.ID,
		URL:        core.URL,
		Name:       core.Name,
		StoreName:  core.StoreName,
		Price:      core.Price,
		City:       core.City,
		MeanRating: core.MeanRating,
		Detail:     core.Detail,
	}
}

//	GetProductRating
func FromCoreRating(core product.CoreRating) ProductRating {
	return ProductRating{
		ID:          core.ID,
		URL:         core.URL,
		ProductName: core.ProductName,
		Rating:      core.Rating,
		Review:      core.Review,
	}
}

func FromCoreListRating(data []product.CoreRating) []ProductRating {
	res := []ProductRating{}
	for key := range data {
		res = append(res, FromCoreRating(data[key]))
	}
	return res
}

//	GetMyProduct
func FromCoreMyProduct(core product.Core) MyProduct {
	return MyProduct{
		ID:    core.ID,
		URL:   core.URL,
		Name:  core.Name,
		Price: core.Price,
		City:  core.City,
	}
}

func FromCoreListMyProduct(data []product.Core) []MyProduct {
	result := []MyProduct{}
	for key := range data {
		result = append(result, FromCoreMyProduct(data[key]))
	}
	return result
}
