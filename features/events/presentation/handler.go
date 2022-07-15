package presentation

import (
	"fmt"
	"lami/app/config"
	"lami/app/features/events"
	"log"
	"net/http"
	"time"

	_request_event "lami/app/features/events/presentation/request"
	_response_event "lami/app/features/events/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventBusiness events.Business
}

func NewEventHandler(business events.Business) *EventHandler {
	return &EventHandler{
		eventBusiness: business,
	}
}

func (h *EventHandler) GetAll(c echo.Context) error {
	name := c.QueryParam("name")
	city := c.QueryParam("city")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	result, total, err := h.eventBusiness.GetAllEvent(limit, pageint, name, city)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get all data"))
	}
	respons := _response_event.FromCoreList(result)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithDataPage("Success get all events", total, respons))
}

func (h *EventHandler) GetDataById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("eventID"))
	result, err := h.eventBusiness.GetEventByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get event"))
	}

	response := _response_event.FromCoreByID(result)

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("success get event", response))
}

func (h *EventHandler) InsertData(c echo.Context) error {
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed insert data"))
	}

	event := _request_event.Event{}
	err_bind := c.Bind(&event)

	if err_bind != nil {
		log.Print(err_bind)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed bind data"))
	}

	layout_time := "2006-01-02T15:04"
	DateTime, errDate := time.Parse(layout_time, event.Date)
	if errDate != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed format date"))
	}
	event.DateTime = DateTime

	imageData, imageInfo, imageErr := c.Request().FormFile("image")
	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get image"))
	}

	imageExtension, err_image_extension := helper.CheckImageExtension(imageInfo.Filename)
	if err_image_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("image extension error"))
	}

	fileData, fileInfo, fileErr := c.Request().FormFile("document")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get file"))
	}

	fileExtension, err_file_extension := helper.CheckFileExtension(fileInfo.Filename)
	if err_file_extension != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file extension error"))
	}

	// check image size
	err_image_size := helper.CheckImageSize(imageInfo.Size)
	if err_image_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file size error"))
	}

	// check file size
	err_file_size := helper.CheckFileSize(fileInfo.Size)
	if err_file_size != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("file size error"))
	}

	// memberikan nama file
	imageName := strconv.Itoa(userID_token) + "_" + event.Name + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	fileName := strconv.Itoa(userID_token) + "_" + event.Name + time.Now().Format("2006-01-02 15:04:05") + "." + fileExtension

	image, errUploadImg := helper.UploadFileToS3(config.EventImages, imageName, config.Images, imageData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to upload file"))
	}

	file, errUploadFile := helper.UploadFileToS3(config.EventImages, fileName, config.Documents, fileData)

	if errUploadFile != nil {
		fmt.Println(errUploadFile)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to upload file"))
	}

	eventCore := _request_event.ToCore(event)
	eventCore.IDUser = userID_token
	eventCore.Image = image
	eventCore.Document = file

	err := h.eventBusiness.InsertEvent(eventCore)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed insert event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessCreate("success insert event"))

}

func (h *EventHandler) DeleteData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("eventID"))

	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get user id"))
	}

	err := h.eventBusiness.DeleteEventByID(id, userID_token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed delete event"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessDelete("success delete data"))
}

func (h *EventHandler) UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("eventID"))
  status := c.Param("status")
	eventReq := _request_event.Event{}
	err_bind := c.Bind(&eventReq)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to bind data"))
	}
	
	userID_token, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get user id"))
	}

	eventCore := _request_event.ToCore(eventReq)
  eventCore.Status = status

	err := h.eventBusiness.UpdateEventByID(eventCore, id, userID_token)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed update data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("sucsess update event"))
}

func (h *EventHandler) GetEventByUser(c echo.Context) error {

	page := c.QueryParam("page")
	pageint, _ := strconv.Atoi(page)
	limitint := 12

	id_user, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get user id"))
	}

	result, total, err := h.eventBusiness.GetEventByUserID(id_user, limitint, pageint)

	respons := _response_event.FromCoreList(result)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get all my events"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithDataPage("success get all my events", total, respons))
}
