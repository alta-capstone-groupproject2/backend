package request

import "lami/app/features/orders"

type Order struct {
	CartID      []int   `json:"cart_id" form:"cart_id"`
	UserID      int     `json:"user_id" form:"user_id"`
	Receiver    string  `json:"receiver" form:"receiver"`
	PhoneNumber string  `json:"phone" form:"phone"`
	Address     string  `json:"address" form:"address"`
	TotalPrice  float32 `json:"totalprice" form:"totalprice"`
	Status      string  `json:"status" form:"status"`
	// Payment     string  `json:"payment" form:"payment"`
}

type OrderDetail struct {
	OrderID     int     `json:"order_id" form:"order_id"`
	Price       float32 `json:"price" form:"price"`
	Qty         uint    `json:"qty" form:"qty"`
}

func ToCore(orderReq Order) orders.Core {
	var dataCartID []int
	for key := range orderReq.CartID {
		dataCartID = append(dataCartID, key)
	}

	return orders.Core{
		UserID:      orderReq.UserID,
		CartID:      dataCartID,
		Receiver:    orderReq.Receiver,
		PhoneNumber: orderReq.PhoneNumber,
		Address:     orderReq.Address,
	}

}

// func ToCoreDetail(orderDetailReq OrderDetail) orders.CoreDetail {
// 	return orders.CoreDetail{
// 		OrderID:     orderDetailReq.OrderID,
// 		ProductName: orderDetailReq.ProductName,
// 		Price:       orderDetailReq.TotalPrice,
// 		Qty:         orderDetailReq.Qty,
// 	}
// }
