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

	return c.JSON(helper.ResponseStatusOkNoData("Success to insert culture"))
}

func (h *CultureHandler) GetCulture(c echo.Context) error {
	limit := c.QueryParam("limit")
	page := c.QueryParam("page")
	limitint, errLimit := strconv.Atoi(limit)
	pageint, errPage := strconv.Atoi(page)

	if errLimit != nil || errPage != nil || limitint == 0 || pageint == 0 {
		return c.JSON(helper.ResponseBadRequest("wrong query param"))
	}

	res, total, err := h.cultureBusiness.SelectCulture(limitint, pageint)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	resp := response.FromCoreList(res)
	return c.JSON(helper.ResponseStatusOkWithDataPage("Success get all your cultures", total, resp))
}

func (h *CultureHandler) GetCulturebyIDCulture(c echo.Context) error {

	cultureID, errConv := strconv.Atoi(c.Param("cultureID"))
	if errConv != nil {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}

	res, err := h.cultureBusiness.SelectCulturebyCultureID(cultureID)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("Failed get culture by cultureID"))
	}

	response := response.FromCore(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get culture by cultureID", response))
}

func (h *CultureHandler) PutCulture(c echo.Context) error {
	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}

	cultureID, errConv := strconv.Atoi(c.Param("cultureID"))
	if errConv != nil {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}

	culture := request.Culture{}
	err_bind := c.Bind(&culture)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data update culture"))
	}

	cultureCore := request.ToCoreUpdate(culture)

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(helper.ResponseBadRequest("failed to get file"))
		}
	}

	err := h.cultureBusiness.UpdateCulture(cultureCore, cultureID, fileInfo, fileData)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("Failed to update data culture"))
	}
	return c.JSON(helper.ResponseStatusOkNoData("Success to update data culture"))

}

func (h *CultureHandler) DeleteCulture(c echo.Context) error {
	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed get id user"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}

	cultureID, errConv := strconv.Atoi(c.Param("cultureID"))
	if errConv != nil {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}

	err := h.cultureBusiness.DeleteCulture(cultureID)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("Failed to delete data culture"))
	}

	return c.JSON(helper.ResponseStatusOkNoData("Success to delete data culture"))
}

func (h *CultureHandler) PostCultureReport(c echo.Context) error {
	userIDToken, _, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to extract token"))
	}

	cultureID, errConv := strconv.Atoi(c.Param("cultureID"))
	if errConv != nil {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}

	report := request.Report{}
	err_bind := c.Bind(&report)

	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data report"))
	}

	reportCore := request.ToCoreReport(report)
	reportCore.CultureID = cultureID
	reportCore.UserID = userIDToken

	err := h.cultureBusiness.AddCultureReport(reportCore)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("Failed to insert culture report"))
	}

	return c.JSON(helper.ResponseStatusOkNoData("Success to insert culture report"))

}

func (h *CultureHandler) GetCultureReport(c echo.Context) error {
	userIDToken, userRole, errToken := middlewares.ExtractToken(c)
	if userIDToken == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed get id user"))
	}
	if userRole != config.Admin {
		return c.JSON(helper.ResponseForbidden("role not authorized"))
	}

	cultureID, errConv := strconv.Atoi(c.Param("cultureID"))
	if errConv != nil {
		return c.JSON(helper.ResponseBadRequest("wrong param"))
	}

	res, err := h.cultureBusiness.SelectReport(cultureID)
	if err != nil {
		return c.JSON(helper.ResponseNotFound("Failed get culture report"))
	}

	response := response.FromCoreListReport(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get culture report", response))
}
