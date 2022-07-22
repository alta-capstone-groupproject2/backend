package presentation

import (
	"lami/app/config"
	product "lami/app/features/products"
	"lami/app/features/products/presentation/request"
	"lami/app/features/products/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBusiness product.Business
}

func NewProductHandler(business product.Business) *ProductHandler {
	return &ProductHandler{
		productBusiness: business,
	}
}

func (h *ProductHandler) PostProduct(c echo.Context) error {

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	product := request.Product{}
	err_bind := c.Bind(&product)

	if err_bind != nil {
		return c.JSON(helper.ResponseForbidden("failed to bind data insert product"))
	}

	// layout_time := "2006-01-02T15:04"
	// Time, errDate := time.Parse(layout_time, time.Now().Format(layout_time))
	// fmt.Println("Time:", Time)
	// if errDate != nil {
	// 	return c.JSON(helper.ResponseBadRequest("failed format date"))
	// }

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get file"))
	}

	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename, config.ContentImage)
	if err_check_extension != nil {
		return c.JSON(helper.ResponseBadRequest("file extension error"))
	}

	// check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size, config.ContentImage)
	if err_check_size != nil {
		return c.JSON(helper.ResponseBadRequest("file size error"))
	}

	// memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + product.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := helper.UploadFileToS3(config.ProductImages, fileName, config.ContentImage,  fileData)

	if errUploadImg != nil {
		return c.JSON(helper.ResponseInternalServerError("failed to upload file"))
	}

	productCore := request.ToCore(product)
	productCore.UserID = userID_token
	productCore.URL = url

	row, err := h.productBusiness.AddProduct(productCore)
	if err != nil || row != 1 {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("success to insert product"))

}

func (h *ProductHandler) PutProduct(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))
	product := request.Product{}
	err_bind := c.Bind(&product)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data update product"))
	}

	// layout_time := "2006-01-02T15:04"
	// UpdatedTime, errDate := time.Parse(layout_time, product.Date)
	// if errDate != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed format date"))
	// }

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	productCore := request.ToCoreUpdate(product)

	fileData, fileInfo, fileErr := c.Request().FormFile("image")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(helper.ResponseBadRequest("failed to get file"))
		}

		extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename, config.ContentImage)
		if err_check_extension != nil {
			return c.JSON(helper.ResponseBadRequest("file extension error"))
		}

		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size, config.ContentImage)
		if err_check_size != nil {
			return c.JSON(helper.ResponseBadRequest("file size error"))
		}

		// memberikan nama file
		fileName := strconv.Itoa(userID_token) + "_" + product.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

		url, errUploadImg := helper.UploadFileToS3(config.ProductImages, fileName, config.ContentImage, fileData)

		if errUploadImg != nil {
			return c.JSON(helper.ResponseBadRequest("failed to upload file"))
		}

		productCore.URL = url
	}

	err := h.productBusiness.UpdateProduct(productCore, idProduct, userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}
	return c.JSON(helper.ResponseStatusOkNoData("Success to update data product"))

}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("productID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	err := h.productBusiness.DeleteProduct(idProduct, userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseStatusOkNoData("Success to delete data product"))
}

func (h *ProductHandler) GetProductList(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	pageint, errPage := strconv.Atoi(page)
	limitint, errLimit := strconv.Atoi(limit)
	if errPage != nil || errLimit != nil{
		return c.JSON(helper.ResponseBadRequest("wrong query param"))
	}

	res, total, err := h.productBusiness.SelectProductList(limitint, pageint, name, city)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("failed get all data"))
	}
	resp := response.FromCoreListProductList(res)
	return c.JSON(helper.ResponseStatusOkWithDataPage("Success get all products", total, resp))
}

func (h *ProductHandler) PostProductRating(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	rating := request.Rating{}
	err_bind := c.Bind(&rating)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("failed to bind data product rating"))
	}

	// layout_time := "2006-01-02T15:04"
	// Time, errDate := time.Parse(layout_time, rating.Date)
	// if errDate != nil {
	// 	return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed format date"))
	// }

	ratingCore := request.ToCoreRating(rating)
	ratingCore.UserID = userID_token
	ratingCore.ProductID = idProduct

	row, err := h.productBusiness.AddProductRating(ratingCore)
	if err != nil || row != 1 {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	return c.JSON(helper.ResponseCreateSuccess("Success to insert product rating"))

}

func (h *ProductHandler) GetProductbyIDProduct(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))

	res, err := h.productBusiness.SelectProductbyIDProduct(idProduct)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	response := response.FromCorebyIDProduct(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get product by idProduct", response))
}

func (h *ProductHandler) GetMyProduct(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseForbidden("user not found"))
	}

	res, err := h.productBusiness.SelectMyProduct(userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	resp := response.FromCoreListMyProduct(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get all your products", resp))
}

func (h *ProductHandler) GetProductRating(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("productID"))

	res, err := h.productBusiness.SelectRating(idProduct)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest(err.Error()))
	}

	response := response.FromCoreListRating(res)
	return c.JSON(helper.ResponseStatusOkWithData("Success get product rating", response))
}
