package data

import (
	cart "lami/app/features/carts"
	_dataProduct "lami/app/features/products/data"
	_dataUser "lami/app/features/users/data"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    int                  `json:"user_id" form:"user_id"`
	ProductID int                  `json:"product_id" form:"product_id"`
	Qty       int                  `json:"qty" form:"qty"`
	Product   _dataProduct.Product `gorm:"foreignKey:ProductID"`
	User      _dataUser.User       `gorm:"foreignKey:UserID"`
}

func fromCore(core cart.Core) Cart {
	return Cart{
		UserID:    core.UserID,
		ProductID: core.ProductID,
		Qty:       core.Qty,
	}
}

// For Get Cart
func (data *Cart) toCore() cart.Core {
	return cart.Core{
		ID:        int(data.ID),
		ProductID: data.ProductID,
		Product: _dataProduct.Product{
			Name:  data.Product.Name,
			URL:   data.Product.URL,
			Price: data.Product.Price,
		},
		Qty: data.Qty,
	}
}

func ToCoreList(data []Cart) []cart.Core {
	res := []cart.Core{}
	for key := range data {
		res = append(res, data[key].toCore())
	}
	return res
}
