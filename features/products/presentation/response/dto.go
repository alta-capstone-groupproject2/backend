package response

import (
	"lami/app/features/products"
)

type Product struct {
	ID         int     `json:"product_id" form:"product_id"`
	URL        string  `json:"url" form:"url"`
	Name       string  `json:"name" form:"name"`
	StoreName  string  `json:"store_name" form:"store_name"`
	Price      float32 `json:"price" form:"price"`
	City       string  `json:"city" form:"city"`
	MeanRating float32 `json:"mean_rating" form:"mean_rating"`
	Detail     string  `json:"detail" form:"detail"`
}

type MyProduct struct {
	ID    int     `json:"product_id" form:"product_id"`
	URL   string  `json:"url" form:"url"`
	Name  string  `json:"name" form:"name"`
	Price float32 `json:"price" form:"price"`
	City  string  `json:"city" form:"city"`
}

type ProductRating struct {
	ID          int    `json:"rating_id" form:"rating_id"`
	URL         string `json:"url" form:"url"`
	ProductName string `json:"name" form:"name"`
	Rating      uint   `json:"rating" form:"rating"`
	Review      string `json:"review" form:"review"`
}

func FromCoreProductList(data product.Core) Product {
	return Product{
		ID:    data.ID,
		URL:   data.URL,
		Name:  data.Name,
		Price: data.Price,
		City:  data.City,
	}
}

func FromCoreListProductList(data []product.Core) []Product {
	res := []Product{}
	for key := range data {
		res = append(res, FromCoreProductList(data[key]))
	}
	return res
}

func FromCore(core product.Core) Product {
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

// Get MyProduct
func FromCoreMyProduct(core product.Core) MyProduct {
	return MyProduct{
		ID:    core.ID,
		URL:   core.URL,
		Name:  core.Name,
		Price: core.Price,
		City:  core.City,
	}
}

func FromCoreList(data []product.Core) []MyProduct {
	result := []MyProduct{}
	for key := range data {
		result = append(result, FromCoreMyProduct(data[key]))
	}
	return result
}
