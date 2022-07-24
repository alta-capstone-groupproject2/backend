package request

import (
	cart "lami/app/features/carts"
	_dataProduct "lami/app/features/products/data"
	_dataUser "lami/app/features/users/data"
)

type Cart struct {
	UserID    int
	ProductID int `json:"productID" form:"productID"`
	Qty       int `json:"qty" form:"qty"`
	Product   _dataProduct.Product
	User      _dataUser.User
}

// For Add Cart
func ToCore(cartReq Cart) cart.Core {
	return cart.Core{
		UserID:    cartReq.UserID,
		ProductID: cartReq.ProductID,
		Qty:       cartReq.Qty,
	}
}

// For Update Cart
func ToCoreUpdate(cartReq Cart) cart.Core {
	return cart.Core{
		Qty: cartReq.Qty,
	}
}
