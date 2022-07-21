package data

import (
	_dataCart "lami/app/features/carts/data"
	"lami/app/features/orders"
	_dataProduct "lami/app/features/products/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int                  `json:"user_id" form:"user_id"`
	Receiver    string               `json:"receiver" form:"receiver"`
	PhoneNumber string               `json:"phone" form:"phone"`
	Address     string               `json:"address" form:"address"`
	TotalPrice  uint                 `json:"totalprice" form:"totalprice"`
	Status      string               `json:"status" form:"status"`
	Cart        _dataCart.Cart       `gorm:"foreignKey:UserID"`
	Product     _dataProduct.Product `gorm:"foreignKey:ID"`
	// Payment     string  `json:"payment" form:"payment"`
}

type OrderDetail struct {
	gorm.Model
	OrderID    int   `json:"order_id" form:"order_id"`
	ProductID  int   `json:"product_id" form:"product_id"`
	TotalPrice uint  `json:"price" form:"price"`
	Qty        uint  `json:"qty" form:"qty"`
	Order      Order `gorm:"foreignKey:ID"`
}

type Product struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
	Qty  uint   `json:"qty" form:"qty"`
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
		OrderID:    core.OrderID,
		ProductID:  core.ProductID,
		TotalPrice: core.TotalPrice,
		Qty:        core.Qty,
	}
}

func (data *OrderDetail) toCore() orders.CoreDetail {
	return orders.CoreDetail{
		Receiver:   data.Order.Receiver,
		Address:    data.Order.Address,
		Status:     data.Order.Status,
		TotalPrice: data.TotalPrice,
		// Product:    []orders.Product{},
	}
}

func toCore(data OrderDetail) orders.CoreDetail {
	return data.toCore()
}

func (data *Product) toProductCore() orders.Product {
	return orders.Product{
		ID:   data.ID,
		Name: data.Name,
		Url:  data.Url,
		Qty:  data.Qty,
	}
}

func ToProductCoreList(data []Product) []orders.Product {
	result := []orders.Product{}
	for key := range data {
		result = append(result, data[key].toProductCore())
	}
	return result
}
