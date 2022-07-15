package response

import "lami/app/features/carts"

type Cart struct {
	ID          int     `json:"id" form:"id"`
	Qty         int     `json:"qty" form:"qty"`
	URL         string  `json:"url" form:"url"`
	ProductName string  `json:"product_name" form:"product_name"`
	Price       float32 `json:"price" form:"price"`
}

// For Get Cart
func FromCore(core cart.Core) Cart {
	return Cart{
		ID:          core.ID,
		URL:         core.Product.URL,
		ProductName: core.Product.Name,
		Price:       core.Product.Price,
		Qty:         core.Qty,
	}
}

func FromCoreList(data []cart.Core) []Cart {
	res := []Cart{}
	for key := range data {
		res = append(res, FromCore(data[key]))
	}
	return res
}
