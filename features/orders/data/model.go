package data

import (
	"lami/app/features/orders"
	_dataCart "lami/app/features/carts/data"
	_dataProduct "lami/app/features/products/data"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      int                  `json:"user_id" form:"user_id"`
	Receiver    string               `json:"receiver" form:"receiver"`
	PhoneNumber string               `json:"phone" form:"phone"`
	Address     string               `json:"address" form:"address"`
	Price       float32              `json:"price" form:"price"`
	Status      string               `json:"status" form:"status"`
	Product     _dataProduct.Product `gorm:"foreignKey:UserID"`
	Cart        _dataCart.Cart
	// Payment     string  `json:"payment" form:"payment"`
}

type OrderDetail struct {
	gorm.Model
	OrderID     int     `json:"order_id" form:"order_id"`
	ProductName string  `json:"product_name" form:"product_name"`
	Price       float32 `json:"price" form:"price"`
	Qty         uint    `json:"qty" form:"qty"`
}

type CartCheck struct {
	Cart _dataCart.Cart
}

func fromCore(core orders.Core) Order {
	return Order{
		UserID:      core.UserID,
		Receiver:    core.Receiver,
		PhoneNumber: core.PhoneNumber,
		Address:     core.Address,
		Status:      core.Status,
	}
}
