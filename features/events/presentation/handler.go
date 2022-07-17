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

	result, total, err := h.eventBusiness.GetAllEvent(limit, page, name, city)
	if err != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get all data"))
	}
	respons := _response_event.FromCoreList(result)
	return c.JSON(helper.ResponseStatusOkWithDataPage("success get all events", total, respons))
}

func (h *EventHandler) GetDataById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("eventID"))
	result, err := h.eventBusiness.GetEventByID(id)
	if err != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get event"))
	}

	response := _response_event.FromCoreByID(result)

	return c.JSON(helper.ResponseStatusOkWithData("success get event", response))
}

func (h *EventHandler) InsertData(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseInternalServerError("failed insert data"))
	}

	event := _request_event.Event{}
	err_bind := c.Bind(&event)

	if err_bind != nil {
		log.Print(err_bind)
		return c.JSON(helper.ResponseBadRequest("failed bind data"))
	}

	layout_time := "2006-01-02T15:04"
	DateTime, errDate := time.Parse(layout_time, event.Date)
	if errDate != nil {
		return c.JSON(helper.ResponseInternalServerError("failed format date"))
	}
	event.DateTime = DateTime

	//upload file Image
	imageData, imageInfo, imageErr := c.Request().FormFile("image")
	if imageErr == http.ErrMissingFile || imageErr != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get image"))
	}

	imageExtension, err_image_extension := helper.CheckFileExtension(imageInfo.Filename, config.ContentImage)
	if err_image_extension != nil {
		return c.JSON(helper.ResponseBadRequest("image extension error"))
	}

	// check image size
	err_image_size := helper.CheckFileSize(imageInfo.Size, config.ContentImage)
	if err_image_size != nil {
		return c.JSON(helper.ResponseBadRequest("image size error"))
	}

	// memberikan nama file
	imageName := strconv.Itoa(userID_token) + "_" + event.Name + time.Now().Format("2006-01-02 15:04:05") + "." + imageExtension

	image, errUploadImg := helper.UploadFileToS3(config.EventImages, config.ContentImage, imageName, imageData)

	if errUploadImg != nil {
		fmt.Println(errUploadImg)
		return c.JSON(helper.ResponseInternalServerError("failed to upload file"))
	}

	//upload file PDF
	fileData, fileInfo, fileErr := c.Request().FormFile("document")
	if fileErr == http.ErrMissingFile || fileErr != nil {
		return c.JSON(helper.ResponseInternalServerError("failed to get file"))
	}

	fileExtension, err_file_extension := helper.CheckFileExtension(fileInfo.Filename, config.ContentDocuments)
	if err_file_extension != nil {
		return c.JSON(helper.ResponseBadRequest("file extension error"))
	}

	// check file size
	err_file_size := helper.CheckFileSize(fileInfo.Size, config.ContentDocuments)
	if err_file_size != nil {
		return c.JSON(helper.ResponseBadRequest("file size error"))
	}

	// memberikan nama file
	fileName := strconv.Itoa(userID_token) + "_" + event.Name + time.Now().Format("2006-01-02 15:04:05") + "." + fileExtension

	file, errUploadFile := helper.UploadFileToS3(config.EventDocuments, fileName, config.ContentDocuments, fileData)

	if errUploadFile != nil {
		fmt.Println(errUploadFile)
		return c.JSON(helper.ResponseInternalServerError("failed to upload file"))
	}

	eventCore := _request_event.ToCore(event)
	eventCore.UserID = userID_token
	eventCore.Image = image
	eventCore.Document = file
	eventCore.Status = config.Status

	err := h.eventBusiness.InsertEvent(eventCore)
	if err != nil {
		fmt.Println(err)
		return c.JSON(helper.ResponseInternalServerError("failed insert event"))
	}
	return c.JSON(helper.ResponseCreateSuccess("success insert event"))

}

func (h *EventHandler) DeleteData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("eventID"))

	userID_token, role, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get user id"))
	}
	if role == config.Admin {
		err := h.eventBusiness.DeleteEventByID(id, userID_token)
		if err != nil {
			return c.JSON(helper.ResponseInternalServerError("failed delete event"))
		}
	}
	return c.JSON(helper.ResponseNoContent("success delete data"))
}

func (h *EventHandler) UpdateData(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	statusReq := _request_event.UpdateEvent{}
	err_bind := c.Bind(&statusReq)

	if err_bind != nil {
		log.Print(err_bind)
		return c.JSON(helper.ResponseBadRequest("failed bind data"))
	}
	status := statusReq.Status
	userID_token, role, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseInternalServerError("failed to get user id"))
	}

	if role == config.Admin {
		err := h.eventBusiness.UpdateEventByID(status, id)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(helper.ResponseInternalServerError("failed update data"))
		}
		return c.JSON(helper.ResponseStatusOkNoData("sucsess update event"))
	}
	return c.JSON(helper.ResponseBadRequest("only is admin"))
}

func (h *EventHandler) GetEventByUser(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	idUser, _, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get user id"))
	}

	result, total, err := h.eventBusiness.GetEventByUserID(idUser, limit, page)

	respons := _response_event.FromCoreList(result)
	if err != nil {
		return c.JSON(helper.ResponseInternalServerError("failed to get all my events"))
	}
	return c.JSON(helper.ResponseStatusOkWithDataPage("success get all my events", total, respons))
}

func (h *EventHandler) GetSubmissionAll(c echo.Context) (err error) {
	page, _ := strconv.Atoi(c.Param("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	_, role, errToken := middlewares.ExtractToken(c)
	if errToken != nil {
		return c.JSON(helper.ResponseInternalServerError("failed get user id"))
	}

	if role == config.Admin {
		result, total, err := h.eventBusiness.GetEventSubmission(limit, page)
		if err != nil {
			return c.JSON(helper.ResponseInternalServerError("failed to get all apply events"))
		}
		response := _response_event.FromSubmissionCoreList(result)
		return c.JSON(helper.ResponseStatusOkWithDataPage("success to get all apply events", total, response))
	}
	return c.JSON(helper.ResponseBadRequest("only is admin"))
}
