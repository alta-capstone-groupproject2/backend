package business

import (
	"errors"
	"lami/app/features/participants"
)

type participantUseCase struct {
	participantData participants.Data
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
func (uc *participantUseCase) AddParticipant(partRequest participants.Core) error {
	if partRequest.EventID == 0 {
		return errors.New("data must be filled")
	}

	err := uc.participantData.AddData(partRequest)
	return err
}

func NewParticipantBusiness(ptrData participants.Data) participants.Business {
	return &participantUseCase{
		participantData: ptrData,
	}
}
