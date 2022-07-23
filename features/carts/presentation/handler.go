package presentation

import (
	cart "lami/app/features/carts"
	"lami/app/features/carts/presentation/request"
	"lami/app/features/carts/presentation/response"

	"lami/app/helper"
	"lami/app/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness cart.Business
}

func NewCartHandler(business cart.Business) *CartHandler {
	return &CartHandler{
		cartBusiness: business,
	}
}

func (h *CartHandler) PostCart(c echo.Context) error {

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	cart := request.Cart{}
	err_bind := c.Bind(&cart)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data cart"))
	}

	cartCore := request.ToCore(cart)
	cartCore.UserID = userID_token
	cartCore.Qty = 1

	row, err := h.cartBusiness.AddCart(cartCore)
	if err != nil || row != 0 {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("Success to insert cart"))

}

func (h *CartHandler) GetCart(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	res, err := h.cartBusiness.SelectCart(userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	resp := response.FromCoreList(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get all your carts", resp))
}

func (h *CartHandler) PutCart(c echo.Context) error {

	idCart, _ := strconv.Atoi(c.Param("cartID"))
	cart := request.Cart{}
	err_bind := c.Bind(&cart)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data update cart"))
	}

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	cartCore := request.ToCoreUpdate(cart)
	err := h.cartBusiness.UpdateCart(cartCore, idCart)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseStatusOkNoData("Success to update data cart"))
}

func (h *CartHandler) DeletedCart(c echo.Context) error {
	idCart, _ := strconv.Atoi(c.Param("cartID"))
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	err := h.cartBusiness.DeleteCart(idCart, userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseNoContent("Success to delete data cart"))

}
