package presentation

import (
	"fmt"
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
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to extract token"))
	}

	product := request.Product{}
	err_bind := c.Bind(&product)

	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to bind data product"))
	}

	layout_time := "2006-01-02T15:04"
	DateTime, errDate := time.Parse(layout_time, product.Date)
	if errDate != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to format date"))
	}
	product.DateTime = DateTime

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get file"))
	}

	extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename, config.ContentImage)
	if err_check_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file extension error"))
	}

	// check file size
	err_check_size := helper.CheckFileSize(fileInfo.Size, config.ContentImage)
	if err_check_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file size error"))
	}

	// memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + product.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

	url, errUploadImg := helper.UploadFileToS3(config.ProductImages, config.ContentImage, fileName, fileData)

	if errUploadImg != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to upload file"))
	}

	productCore := request.ToCore(product)
	productCore.UserID = userID_token
	productCore.URL = url

	row, err := h.productBusiness.AddProduct(productCore)
	fmt.Println(row)
	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed to insert product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to insert product"))

}

func (h *ProductHandler) PutProduct(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))
	product := request.Product{}
	err_bind := c.Bind(&product)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to bind data update product"))
	}

	layout_time := "2006-01-02T15:04"
	DateTime, errDate := time.Parse(layout_time, product.Date)
	if errDate != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed format date"))
	}
	product.DateTime = DateTime

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get id user"))
	}

	productCore := request.ToCoreUpdate(product)

	fileData, fileInfo, fileErr := c.Request().FormFile("file")
	if fileErr != http.ErrMissingFile {
		if fileErr != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get file"))
		}

		extension, err_check_extension := helper.CheckFileExtension(fileInfo.Filename, config.ContentImage)
		if err_check_extension != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file extension error"))
		}

		// check file size
		err_check_size := helper.CheckFileSize(fileInfo.Size, config.ContentImage)
		if err_check_size != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file size error"))
		}

		// memberikan nama file
		fileName := strconv.Itoa(userID_token) + "_" + product.Name + time.Now().Format("2006-01-02 15:04:05") + "." + extension

		url, errUploadImg := helper.UploadFileToS3(config.ProductImages, config.ContentImage, fileName, fileData)

		if errUploadImg != nil {
			fmt.Println(errUploadImg)
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to upload file"))
		}

		productCore.URL = url
	}

	err := h.productBusiness.UpdateProduct(productCore, idProduct, userID_token)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed to update data product"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to update data product"))

}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("productID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get id user"))
	}

	err := h.productBusiness.DeleteProduct(idProduct, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed to delete data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to delete data product"))
}

func (h *ProductHandler) PostProductRating(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to extract token"))
	}

	rating := request.Rating{}
	err_bind := c.Bind(&rating)

	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to bind data rating"))
	}

	ratingCore := request.ToCoreRating(rating)
	ratingCore.UserID = userID_token
	ratingCore.ProductID = idProduct

	row, err := h.productBusiness.AddProductRating(ratingCore)
	fmt.Println(row)
	if err != nil || row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed to insert product rating"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to insert product rating"))

}

func (h *ProductHandler) GetProductbyIDProduct(c echo.Context) error {

	idProduct, _ := strconv.Atoi(c.Param("productID"))

	res, err := h.productBusiness.SelectProductbyIDProduct(idProduct)
	fmt.Println("res from handler.go:", res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed get product by idProduct"))
	}

	response := response.FromCore(res)
	fmt.Println("response from handler.go:", response)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get product by idProduct", response))
}

func (h *ProductHandler) GetMyProduct(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to extract token"))
	}

	res, err := h.productBusiness.SelectMyProduct(userID_token)
	fmt.Println("res from handler:", res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get all your products"))
	}

	resp := response.FromCoreList(res)
	fmt.Println("response:", resp)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get all your products", resp))
}

func (h *ProductHandler) GetProductRating(c echo.Context) error {
	idProduct, _ := strconv.Atoi(c.Param("productID"))

	res, err := h.productBusiness.SelectRating(idProduct)
	fmt.Println("res rating from handler.go:", res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("Failed get product rating"))
	}

	response := response.FromCoreListRating(res)
	fmt.Println("response rating from handler.go:", response)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success get product rating", response))
}
