package business

import (
	"errors"
	"fmt"
	"lami/app/features/participants"
	"lami/app/helper"

	"github.com/midtrans/midtrans-go/coreapi"
)

type participantUseCase struct {
	participantData participants.Data
}

func NewParticipantBusiness(ptrData participants.Data) participants.Business {
	return &participantUseCase{
		participantData: ptrData,
	}
}

// DeleteParticipan implements participants.Business
func (uc *participantUseCase) DeleteParticipant(param, userID int) error {
	err := uc.participantData.DeleteData(param, userID)
	if err != nil {
		return errors.New("no data user for deleted")
	}
	return nil
}

func (uc *participantUseCase) GetAllEventbyParticipant(idUser int) (data []participants.Core, err error) {
	resp, err := uc.participantData.SelectDataEvent(idUser)
	return resp, err
}

// AddParticipant implements participants.Business
func (uc *participantUseCase) AddParticipant(partRequest participants.Core) (err error) {
	if partRequest.EventID == 0 {
		return errors.New("data must be filled")
	}
	/** Check event date by eventID
	** @Param EventID
	** @return eventData struct
	**/
	checkEvent, errCheckEvent := uc.participantData.SelectDataByID(partRequest.EventID)
	if errCheckEvent != nil {
		return errors.New("no data event")
	}
	/** Check Date By Array Participation
	** @Param 	UserID
	** @return 	[]Index data
	** @return 	error
	**/
	checkDate, errCheckDate := uc.participantData.SelectDataEvent(partRequest.UserID)
	if errCheckDate != nil {
		return errors.New("failed get data join event")
	}

	for i := 0; i < len(checkDate); i++ {
		if checkDate[i].Event.Date.Before(checkEvent.EndDate) && checkDate[i].Event.Date.After(checkEvent.StartDate) {
			return errors.New("can't both events")
		}
	}

	result := uc.participantData.SelectValidasi(partRequest.UserID, partRequest.EventID)
	if result {
		return errors.New("user have a join")
	}
	err = uc.participantData.AddData(partRequest)
	if err == nil {
		eventData, _ := uc.participantData.SelectDataByID(partRequest.EventID)
		userData, _ := uc.participantData.SelectUser(partRequest.UserID)
		detailMal := fmt.Sprintf("Event : %s, \nDate : %s", eventData.Name, eventData.StartDate)

		helper.SendGmailNotify(userData.Email, "Join Event", detailMal)
	}

	return err
}

func (uc *participantUseCase) GrossAmountEvent(eventID int) (GrossAmount int64, err error) {
	checkEvent, errCheckEvent := uc.participantData.SelectDataByID(eventID)
	if errCheckEvent != nil {
		return 0, errCheckEvent
	}
	return int64(checkEvent.Price), nil
}

func (uc *participantUseCase) CreatePaymentBankTransfer(reqPay coreapi.ChargeReq, reqJoin participants.Core) (resPay *coreapi.ChargeResponse, err error) {
	createPay, errCreatePay := uc.participantData.CreateDataPayment(reqPay)
	if errCreatePay != nil {
		return nil, errors.New("failed get response payment")
	}
	updateJoin := uc.participantData.UpdateDataPayment(createPay, reqJoin)
	if updateJoin != nil {
		return nil, errors.New("failed to connect update")
	}
	return createPay, nil
}

func (uc *participantUseCase) GetDetailPayment(limit, page, userID int) (res []participants.Core, total int64, err error) {
	offset := limit * (page - 1)
	result, total, err := uc.participantData.SelectPaymentList(limit, offset, userID)
	if err != nil {
		return []participants.Core{}, 0, err
	}
	total = total/int64(limit) + 1
	return result, total, nil
}

func (uc *participantUseCase) CheckStatusPayment(orderID string) (*coreapi.TransactionStatusResponse, error) {
	result, err := uc.participantData.CheckDataStatusPayment(orderID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *participantUseCase) PaymentWebHook(orderID, status string) error {
	payment, errPayment := uc.participantData.SelectPayment(orderID)
	if errPayment != nil {
		return errors.New("failed to get data join")
	}

	if status == "settlement" {
		payment.Status = "paid"
	}
	if status == "cancel" || status == "deny" || status == "expire" {
		payment.Status = "failed"
		payment.PaymentMethod = ""
		payment.TransactionID = ""
	}

	result := uc.participantData.PaymentDataWebHook(payment)
	if result != nil {
		return result
	}
	return nil
}
