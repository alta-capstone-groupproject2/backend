package response

import (
	product "lami/app/features/products"
)

type Product struct {
	ID         int     `json:"productID"`
	URL        string  `json:"image"`
	Name       string  `json:"productName"`
	StoreName  string  `json:"storeName"`
	Price      uint    `json:"price"`
	City       string  `json:"city"`
	MeanRating float64 `json:"meanRating"`
	Detail     string  `json:"details"`
}

type MyProduct struct {
	ID    int    `json:"productID"`
	URL   string `json:"image"`
	Name  string `json:"productName"`
	Price uint   `json:"price"`
	City  string `json:"city"`
}

type ProductRating struct {
	ID          int     `json:"ratingID"`
	URL         string  `json:"image"`
	ProductName string  `json:"name"`
	Rating      float64 `json:"rating"`
	Review      string  `json:"review"`
}

type ProductList struct {
	ID    int    `json:"productIDd"`
	URL   string `json:"image"`
	Name  string `json:"productName"`
	Price uint   `json:"price"`
	City  string `json:"city"`
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
