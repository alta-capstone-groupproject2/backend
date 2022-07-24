package data

import (
	"lami/app/features/orders"
	"lami/app/features/products/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int    `json:"user_id" form:"user_id"`
	Receiver    string `json:"receiver" form:"receiver"`
	PhoneNumber string `json:"phone" form:"phone"`
	Address     string `json:"address" form:"address"`
	TotalPrice  uint   `json:"totalprice" form:"totalprice"`
	Status      string `json:"status" form:"status"`
	OrderDetail []OrderDetail
}

type OrderDetail struct {
	gorm.Model
	OrderID   int          `json:"order_id" form:"order_id"`
	ProductID int          `json:"product_id" form:"product_id"`
	Name      string       `json:"name" form:"name"`
	URL       string       `json:"url" form:"url"`
	Qty       uint         `json:"qty" form:"qty"`
	Product   data.Product `json:"product" form:"product"`
}

func fromCore(core orders.Core) Order {
	return Order{
		UserID:      core.UserID,
		Receiver:    core.Receiver,
		PhoneNumber: core.PhoneNumber,
		Address:     core.Address,
		TotalPrice:  core.TotalPrice,
		Status:      core.Status,
	}
}

func fromCoreDetail(core orders.CoreDetail) OrderDetail {
	return OrderDetail{
		OrderID:   core.OrderID,
		ProductID: core.ProductID,
		Qty:       core.Qty,
	}
}

func (data *Order) toCore() orders.Core {
	return orders.Core{
		Receiver:   data.Receiver,
		TotalPrice: data.TotalPrice,
		Address:    data.Address,
		Status:     data.Status,
		Product:    ToCoreDetailList(data.OrderDetail),
	}
}

func ToCoreList(data []Order) []orders.Core {
	res := []orders.Core{}
	for v := range data {
		res = append(res, data[v].toCore())
	}
	return res
}

func (data *OrderDetail) toCoreDetail() orders.CoreDetail {
	return orders.CoreDetail{
		Product: orders.Product{
			ID:   int(data.Product.ID),
			Name: data.Product.Name,
			Url:  data.Product.URL,
			Qty:  data.Qty,
		},
	}
}

func ToCoreDetailList(data []OrderDetail) []orders.CoreDetail {
	res := []orders.CoreDetail{}
	for v := range data {
		res = append(res, data[v].toCoreDetail())
	}
	return res
}
