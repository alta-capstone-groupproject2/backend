package business

import (
	"errors"
	"lami/app/features/participants"

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
		if checkDate[i].Event.Date == checkEvent.Date {
			return errors.New("can't both events")
		}
	}
	err = uc.participantData.AddData(partRequest)
	return err
}

func (uc *participantUseCase) GrossAmountEvent(userID int) (GrossAmount int64, err error) {
	checkEvent, errCheckEvent := uc.participantData.SelectDataByID(userID)
	if errCheckEvent != nil {
		return 0, errCheckEvent
	}
	return int64(checkEvent.Price), nil
}

func (uc *participantUseCase) CreatePaymentBankTransfer(reqPay coreapi.ChargeReq) (resPay *coreapi.ChargeResponse, err error) {
	createPay, errCreatePay := uc.participantData.CreateDataPayment(reqPay)
	if errCreatePay != nil {
		return nil, errors.New("failed get response payment")
	}
	return createPay, nil
}
