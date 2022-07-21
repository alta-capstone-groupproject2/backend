package response

import "lami/app/features/orders"

type Order struct {
	UserID      int    `json:"user_id" form:"user_id"`
	Receiver    string `json:"receiver" form:"receiver"`
	PhoneNumber string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	TotalPrice  uint   `json:"totalprice" form:"totalprice"`
	Status      string `json:"status" form:"status"`
	// Payment     string  `json:"payment" form:"payment"`
}

type OrderDetail struct {
	Receiver   string `json:"receiver" form:"receiver"`
	Address    string `json:"address" form:"address"`
	TotalPrice uint   `json:"totalprice" form:"totalprice"`
	Status     string `json:"status" form:"status"`
	Product    []Product
	Order      Order
}

type Product struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
	Qty  uint   `json:"qty" form:"qty"`
}

func FromCore(data orders.CoreDetail) OrderDetail {
	return OrderDetail{
		Receiver:   data.Receiver,
		Address:    data.Address,
		TotalPrice: data.TotalPrice,
		Status:     data.Status,
		Product:    FromOrderCoreList(data.Product),
	}
}

func FromOrderCore(data orders.Product) Product {
	return Product{
		ID:   data.ID,
		Name: data.Name,
		Url:  data.Url,
		Qty:  data.Qty,
	}
}

func FromOrderCoreList(data []orders.Product) []Product {
	res := []Product{}
	for key := range data {
		res = append(res, FromOrderCore(data[key]))
	}
	return res
}
