package presentation

import (
	"fmt"
	"lami/app/config"
	"lami/app/features/participants"
	_request_participant "lami/app/features/participants/presentation/request"
	_response_participant "lami/app/features/participants/presentation/response"
	"lami/app/helper"
	"lami/app/middlewares"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type ParticipantHandler struct {
	participantBusiness participants.Business
}

func NewParticipantHandler(business participants.Business) *ParticipantHandler {
	return &ParticipantHandler{
		participantBusiness: business,
	}
}

func (h *ParticipantHandler) Joined(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	participant := _request_participant.Participant{}
	err_bind := c.Bind(&participant)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("error bind data"))
	}

	participantCore := _request_participant.ToCore(participant)
	participantCore.UserID = userID_token
	participantCore.Status = config.PaymentStatus

	err := h.participantBusiness.AddParticipant(participantCore)
	if err != nil {
		return c.JSON(helper.ResponseInternalServerError("failed, can't both join event"))
	}
	return c.JSON(helper.ResponseCreateSuccess("success join"))

}

func (h *ParticipantHandler) GetAllEventParticipant(c echo.Context) error {
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	result, err := h.participantBusiness.GetAllEventbyParticipant(userID_token)
	if err != nil {
		return c.JSON(helper.ResponseBadRequest("failed get all your events"))
	}

	response := _response_participant.FromCoreList(result)
	return c.JSON(helper.ResponseStatusOkWithData("Success get all your events", response))

}

func (h *ParticipantHandler) DeleteEventbyParticipant(c echo.Context) error {
	idParticipant, _ := strconv.Atoi(c.Param("id"))

	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	result := h.participantBusiness.DeleteParticipant(idParticipant, userID_token)
	if result != nil {
		return c.JSON(helper.ResponseBadRequest("failed to delete your event"))
	}
	return c.JSON(helper.ResponseNoContent("success to delete your event"))
}

func (h *ParticipantHandler) CreatePayment(c echo.Context) error {
	currentTime := time.Now()
	userID_token, _, errToken := middlewares.ExtractToken(c)
	if userID_token == 0 || errToken != nil {
		return c.JSON(helper.ResponseBadRequest("failed to get user id"))
	}

	reqPay := _request_participant.Participant{}
	errBind := c.Bind(&reqPay)
	if errBind != nil {
		return c.JSON(helper.ResponseBadRequest("error bind data"))
	}

	reqPayCore := _request_participant.ToCore(reqPay)

	reqPayCore.UserID = userID_token

	grossAmount, _ := h.participantBusiness.GrossAmountEvent(reqPayCore.EventID)
	reqPayCore.GrossAmount = grossAmount
	date := currentTime.Format("2006-01-02")
	timer := currentTime.Format("15:04:05")

	orderIDPay := fmt.Sprintf("Tiket-%d-%s-%s", reqPay.EventID, date, timer)

	reqPayCore.OrderID = orderIDPay

	inputPay := _request_participant.ToCoreMidtrans(reqPayCore)
	if reqPay.PaymentMethod == "BANK_TRANSFER_BCA" {
		reqPayCore.PaymentMethod = "BANK_TRANSFER_BCA"
		inputPay.BankTransfer = &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
		}
	}
	reqCreatePay, errReqCreatePay := h.participantBusiness.CreatePaymentBankTransfer(inputPay, reqPayCore)

	if errReqCreatePay != nil {
		return c.JSON(helper.ResponseBadRequest("failed to payment"))
	}
	result := _response_participant.FromMidtransToPayment(reqCreatePay)
	layout := "2006-01-02 15:04:05"
	trTime, _ := time.Parse(layout, reqCreatePay.TransactionTime)
	result.TransactionTime = trTime
	result.TransactionExpire = trTime.Add(time.Hour * 24)

	return c.JSON(helper.ResponseStatusOkWithData("success create payment", result))
}

func (h *ParticipantHandler) MidtransWebHook(c echo.Context) error {
	midtransRequest := _request_participant.MidtransHookRequest{}
	err_bind := c.Bind(&midtransRequest)
	if err_bind != nil {
		return c.JSON(helper.ResponseBadRequest("error bind data"))
	}

	errUpdateStatusPay := h.participantBusiness.PaymentWebHook(midtransRequest.OrderID, midtransRequest.TransactionStatus)
	if errUpdateStatusPay != nil {
		return c.JSON(helper.ResponseBadRequest("failed to update status payment"))
	}
	return c.JSON(helper.ResponseStatusOkNoData("success to update status payment"))
}
