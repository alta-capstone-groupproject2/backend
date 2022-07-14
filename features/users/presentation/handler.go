package presentation

import (
	"fmt"
	"lami/app/features/users"
	_requestUser "lami/app/features/users/presentation/request"
	_responseUser "lami/app/features/users/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"
	"strconv"
	"time"

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
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	result, err := h.userBusiness.GetDataById(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to get data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessWithData("success", _responseUser.FromCore(result)))
}

func (h *UserHandler) Insert(c echo.Context) error {
	user := _requestUser.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to bind insert data"))
	}

	userCore := _requestUser.ToCore(user)
	_, err := h.userBusiness.InsertData(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed(err.Error()))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success insert data"))
}

func (h *UserHandler) Delete(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}
	_, err := h.userBusiness.DeleteData(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to delete data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success delete data"))
}

func (h *UserHandler) Update(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get user id"))
	}

	userReq := _requestUser.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to bind update data"))
	}

	//	Dari sini
	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get file"))
	}

	//	Check file extension
	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file extension error"))
	}

	//	Check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("file size error"))
	}

	//	Memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + userReq.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := helper.UploadFileToS3(fileName, fileData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to upload file"))
	}

	userCore := _requestUser.ToCore(userReq)
	userCore.Image = url

	_, err := h.userBusiness.UpdateData(userCore, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailed("failed to insert data"))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success update data"))
}
