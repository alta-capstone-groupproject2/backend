package presentation

import (
	"lami/app/config"
	"lami/app/features/users"
	_requestUser "lami/app/features/users/presentation/request"
	_responseUser "lami/app/features/users/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"log"
	"net/http"
	"strconv"

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

func (h *UserHandler) GetDataById(c echo.Context) error {
	userIDToken, _, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	result, err := h.userBusiness.GetDataById(userIDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer(err.Error()))
	}
	if result.Role.RoleName == config.User {
		return c.JSON(helper.ResponseStatusOkWithData("success", _responseUser.FromCore(result)))
	} else {
		return c.JSON(helper.ResponseStatusOkWithData("success", _responseUser.UserStoreFromCore(result)))
	}

}

func (h *UserHandler) Insert(c echo.Context) error {
	user := _requestUser.User{}
	err_bind := c.Bind(&user)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("error bind data"))
	}

	userCore := _requestUser.ToCore(user)
	err := h.userBusiness.InsertData(userCore)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseCreateSuccess("success insert data"))
}

func (h *UserHandler) Delete(c echo.Context) error {
	userIDToken, _, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}
	err := h.userBusiness.DeleteData(userIDToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			helper.ResponseFailedServer("failed to delete data"))
	}
	return c.JSON(helper.ResponseStatusOkNoData("success delete data"))
}

func (h *UserHandler) Update(c echo.Context) error {
	userIDToken, _, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	userReq := _requestUser.User{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind update data"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(helper.ResponseBadRequest("failed to get file"))
		}
	}

	userCore := _requestUser.ToCore(userReq)

	err := h.userBusiness.UpdateData(userCore, userIDToken, fileInfo, fileData)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkNoData("success update data"))
}

func (h *UserHandler) AccountUpgrade(c echo.Context) error {
	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	if userRole == config.UMKM {
		return c.JSON(helper.ResponseBadRequest("your account already upgraded"))
	}

	userReq := _requestUser.Store{}
	err_bind := c.Bind(&userReq)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind update data"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("document")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			log.Print(fileErr)
			return c.JSON(helper.ResponseBadRequest("failed to get file"))
		}
	}

	userCore := _requestUser.StoreToCore(userReq)

	err := h.userBusiness.UpgradeAccount(userCore, userIDToken, fileInfo, fileData)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkNoData("success update data"))
}

func (h *UserHandler) UpdateStatusAccount(c echo.Context) error {
	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("access denied"))
	}

	userId, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil || userIDToken == userId {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}
	dataReq := _requestUser.Store{}
	errBind := c.Bind(&dataReq)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind update data"))
	}

	err := h.userBusiness.UpdateStatusUser(dataReq.StoreStatus, userId)
	if err != nil {
		return c.JSON(helper.ResponseInternalServerError(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkNoData("success update data"))
}

func (h *UserHandler) GetStoreSubmission(c echo.Context) error {

	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	limitint, errLimit := strconv.Atoi(limit)
	pageint, errPage := strconv.Atoi(page)

	if errLimit != nil || errPage != nil || limitint == 0 || pageint == 0 {
		return c.JSON(helper.ResponseBadRequest("wrong query param"))
	}

	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("access denied"))
	}

	result, totalPage, err := h.userBusiness.GetDataSubmissionStore(limitint, pageint)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get all data"))
	}

	return c.JSON(helper.ResponseStatusOkWithDataPage("success", totalPage, _responseUser.UserStoreFromCoreList(result)))
}

func (h *UserHandler) GmailVerification(c echo.Context) error {
	userData := _requestUser.User{}
	errBind := c.Bind(&userData)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("failed bind data"))
	}
	userCore := _requestUser.ToCore(userData)
	errVerify := h.userBusiness.VerifyEmail(userCore)
	if errVerify != nil {
		return c.JSON(helper.ResponseBadRequest(errVerify.Error()))
	}
	return c.JSON(http.StatusOK,
		helper.ResponseSuccessNoData("success email verification sent"))
}

func (h *UserHandler) InsertFromVerificaton(c echo.Context) error {
	encrypt := c.Param("encrypt")

	err := h.userBusiness.ConfirmEmail(encrypt)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseCreateSuccess("success insert data"))
}
