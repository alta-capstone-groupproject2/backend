package presentation

import (
	"lami/app/config"
	culture "lami/app/features/cultures"
	"lami/app/features/cultures/presentation/request"
	"lami/app/features/cultures/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CultureHandler struct {
	cultureBusiness culture.Business
}

func NewCultureHandler(business culture.Business) *CultureHandler {
	return &CultureHandler{
		cultureBusiness: business,
	}
}

func (h *CultureHandler) PostCulture(c echo.Context) error {

	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}

	culture := request.Culture{}
	err_bind := c.Bind(&culture)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data culture"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get file"))
	}

	cultureCore := request.ToCore(culture)

	err := h.cultureBusiness.AddCulture(cultureCore, fileInfo, fileData)

	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to insert culture"))

}

func (h *CultureHandler) PutCulture(c echo.Context) error {

	idCulture, _ := strconv.Atoi(c.Param("cultureID"))
	culture := request.Culture{}
	err_bind := c.Bind(&culture)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data update culture"))
	}

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to get id user"))
	}

	cultureCore := request.ToCoreUpdate(culture)

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(helper.ResponseBadRequest("failed to get file"))
		}
	}

	err := h.cultureBusiness.UpdateCulture(cultureCore, idCulture, fileInfo, fileData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update data culture"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to update data culture"))

}

func (h *CultureHandler) DeleteCulture(c echo.Context) error {
	idCulture, _ := strconv.Atoi(c.Param("cultureID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get id user"))
	}

	err := h.cultureBusiness.DeleteCulture(idCulture, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to delete data culture"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to delete data culture"))
}

func (h *CultureHandler) PostCultureReport(c echo.Context) error {

	idCulture, _ := strconv.Atoi(c.Param("cultureID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to extract token"))
	}

	report := request.Report{}
	err_bind := c.Bind(&report)

	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to bind data report"))
	}

	reportCore := request.ToCoreReport(report)
	reportCore.CultureID = idCulture

	err := h.cultureBusiness.AddCultureReport(reportCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to insert culture report"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to insert culture report"))

}

func (h *CultureHandler) GetCulturebyIDCulture(c echo.Context) error {

	idCulture, _ := strconv.Atoi(c.Param("cultureID"))

	res, err := h.cultureBusiness.SelectCulturebyCultureID(idCulture)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed get culture by idCulture"))
	}

	response := response.FromCore(res)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get culture by idCulture", response))
}

func (h *CultureHandler) GetCulture(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to extract token"))
	}

	res, err := h.cultureBusiness.SelectMyCulture(userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed get all your cultures"))
	}

	resp := response.FromCoreList(res)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get all your cultures", resp))
}

func (h *CultureHandler) GetCultureReport(c echo.Context) error {
	idCulture, _ := strconv.Atoi(c.Param("cultureID"))

	res, err := h.cultureBusiness.SelectReport(idCulture)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed get culture report"))
	}

	response := response.FromCoreListReport(res)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get culture report", response))
}
