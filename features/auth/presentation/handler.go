package presentation

import (
	"lami/app/features/auth"
	"lami/app/features/auth/presentation/request"
	"lami/app/features/auth/presentation/response"
	"lami/app/helper"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userBusiness auth.Business
}

func NewAuthHandler(business auth.Business) *AuthHandler {
	return &AuthHandler{
		userBusiness: business,
	}
}

func (h *AuthHandler) Login(c echo.Context) error {
	reqBody := request.User{}
	errBind := c.Bind(&reqBody)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("failed bind data"))
	}

	authCore := request.ToCore(reqBody)
	token, id, role, err := h.userBusiness.Login(authCore)

	result := response.ToResponse(id, role, token)
	if err != nil {
		return c.JSON(helper.ResponseNotFound(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkWithData("login success", result))
}
