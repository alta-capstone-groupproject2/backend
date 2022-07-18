package business

import (
	"errors"
	"lami/app/features/participants"
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
	** @Param UserID
	** @return []Index data, error
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
