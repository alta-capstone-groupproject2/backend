package presentation

import (
	"lami/app/features/users"
	_requestUser "lami/app/features/users/presentation/request"
	_responseUser "lami/app/features/users/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(business users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: business,
	}
}

// func (h *UserHandler) GetAll(c echo.Context) error {
// 	limit := c.QueryParam("limit")
// 	offset := c.QueryParam("offset")
// 	limitint, _ := strconv.Atoi(limit)
// 	offsetint, _ := strconv.Atoi(offset)
// 	result, err := h.userBusiness.GetAllData(limitint, offsetint)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 			helper.ResponseFailed("failed to get all data"))
// 	}

// 	return c.JSON(http.StatusOK,
// 		helper.ResponseSuccessWithData("success", _responseUser.FromCoreList(result)))
// }

func (h *UserHandler) GetDataById(c echo.Context) error {
	userIDToken, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get user id"))
	}

	result, err := h.userBusiness.GetDataById(userIDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer(err.Error()))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) Insert(c echo.Context) error {
	user := _requestUser.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer("failed to bind insert data"))
	}

	userCore := _requestUser.ToCore(user)
	err := h.userBusiness.InsertData(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer(err.Error()))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success insert data"))
}

func (h *UserHandler) Delete(c echo.Context) error {
	userIDToken, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get user id"))
	}
	err := h.userBusiness.DeleteData(userIDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer("failed to delete data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success delete data"))
}

func (h *UserHandler) Update(c echo.Context) error {
	userIDToken, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get user id"))
	}

	userReq := _requestUser.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer("failed to bind update data"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			log.Print(fileErr)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get file"))
		}
	}

	userCore := _requestUser.ToCore(userReq)

	err := h.userBusiness.UpdateData(userCore, userIDToken, fileInfo, fileData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer(err.Error()))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success update data"))
}
