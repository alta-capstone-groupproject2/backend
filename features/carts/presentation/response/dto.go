package response

import cart "lami/app/features/carts"

type Cart struct {
	ID          int    `json:"cartID"`
	ProductID   int    `json:"productID"`
	Qty         int    `json:"qty"`
	URL         string `json:"image"`
	ProductName string `json:"name"`
	Price       uint   `json:"price"`
}

// For Get Cart
func FromCore(core cart.Core) Cart {
	return Cart{
		ID:          core.ID,
		ProductID:   core.ProductID,
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
