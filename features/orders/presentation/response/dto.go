package response

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
	OrderID     int    `json:"order_id" form:"order_id"`
	ProductName string `json:"product_name" form:"product_name"`
	TotalPrice  uint   `json:"totalprice" form:"totalprice"`
	Qty         uint   `json:"qty" form:"qty"`
}
