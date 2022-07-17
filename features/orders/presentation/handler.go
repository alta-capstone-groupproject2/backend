package presentation

import (
	"lami/app/features/orders"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBusiness orders.Business
}

func NewOrderHandler(business orders.Business) *OrderHandler {
	return &OrderHandler{
		orderBusiness: business,
	}
}

func (h *OrderHandler) PostOrder(c echo.Context) error {
	
}