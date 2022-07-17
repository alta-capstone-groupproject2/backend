package request

type Order struct {
	UserID      int     `json:"user_id" form:"user_id"`
	Receiver    string  `json:"receiver" form:"receiver"`
	PhoneNumber string  `json:"phone" form:"phone"`
	Address     string  `json:"address" form:"address"`
	Price       float32 `json:"price" form:"price"`
	Status      string  `json:"status" form:"status"`
	// Payment     string  `json:"payment" form:"payment"`
}

type OrderDetail struct {
	OrderID     int     `json:"order_id" form:"order_id"`
	ProductName string  `json:"product_name" form:"product_name"`
	Price       float32 `json:"price" form:"price"`
	Qty         uint    `json:"qty" form:"qty"`
}
