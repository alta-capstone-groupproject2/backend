package presentation

import (
	"lami/app/features/comments"
	_request_comment "lami/app/features/comments/presentation/request"
	_response_comment "lami/app/features/comments/presentation/response"
	"net/http"
	"strconv"

	"lami/app/helper"
	"lami/app/middlewares"

	// "strconv"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentBusiness comments.Business
}

func NewCommentHandler(business comments.Business) *CommentHandler {
	return &CommentHandler{
		commentBusiness: business,
	}
}

func (h *CommentHandler) Add(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to get user id"))
	}

	comment := _request_comment.Comment{}
	err_bind := c.Bind(&comment)
	if err_bind != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed to bind insert data"))
	}

	commentCore := _request_comment.ToCore(comment)
	commentCore.UserID = userID_token

	row, err := h.commentBusiness.AddComment(commentCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed insert your comment"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailedBadRequest("failed insert your comment"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessCreate("success insert your comment"))
}

func (h *CommentHandler) Get(c echo.Context) error {
	eventId, _ := strconv.Atoi(c.Param("id"))
	limit := 5
	offset, _ := strconv.Atoi(c.QueryParam("page"))

	result, total, err := h.commentBusiness.GetCommentByIdEvent(limit, offset, eventId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailedServer("failed get all comment events"))
	}
	respons := _response_comment.FromCoreList(result)
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithDataPage("success get all comment events", total, respons))
}
