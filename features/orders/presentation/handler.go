package presentation

import (
	"lami/app/features/orders"
	"lami/app/features/orders/presentation/request"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"

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
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	order := request.Order{}
	err_bind := c.Bind(&order)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data order"))
	}

	orderCore := request.ToCore(order)
	orderCore.UserID = userID_token

	err := h.orderBusiness.AddOrder(orderCore, userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("success to create data order"))
}

func (h *OrderHandler) GetHistoryOrder(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, "failed to extract token")
	}

	panic("unimplemented")
}
