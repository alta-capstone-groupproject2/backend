package bussiness

import (
	"errors"
	"lami/app/features/events"
)

type eventUseCase struct {
	eventData events.Data
}

func NewEventBusiness(usrData events.Data) events.Business {
	return &eventUseCase{
		eventData: usrData,
	}
}

func (uc *eventUseCase) GetAllEvent(limit int, page int, name string, city string) (response []events.Core, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.eventData.SelectData(limit, offset, name, city)
	total = total/int64(limit) + 1
	return resp, total, errData
}

func (uc *eventUseCase) GetEventByID(id int) (response events.Core, err error) {
	response, err = uc.eventData.SelectDataByID(id)
	if err != nil {
		return events.Core{}, err
	}
	responseParticipant, errParticipant := uc.eventData.SelectParticipantData(response.ID)
	response.Participant = responseParticipant
	if errParticipant != nil {
		return events.Core{}, errParticipant
	}
	return response, err
}

func (uc *eventUseCase) InsertEvent(eventRequest events.Core) error {
	if eventRequest.Name == "" || eventRequest.Image == "" || eventRequest.Detail == "" || eventRequest.City == "" || eventRequest.Location == "" || eventRequest.HostedBy == "" {
		return errors.New("all data must be filled")
	}

	err := uc.eventData.InsertData(eventRequest)
	return err
}

func (uc *eventUseCase) DeleteEventByID(id int, userId int) (err error) {
	err = uc.eventData.DeleteDataByID(id, userId)
	return err
}

func (uc *eventUseCase) UpdateEventByID(eventReq events.Core, id int, userId int) (err error) {
	updateMap := make(map[string]interface{})
	if eventReq.Name != "" {
		updateMap["name"] = &eventReq.Name
	}
	if eventReq.Detail != "" {
		updateMap["detail"] = &eventReq.Detail
	}
	if eventReq.City != "" {
		updateMap["city"] = &eventReq.City
	}
	if !eventReq.Date.IsZero() {
		updateMap["date"] = &eventReq.Date
	}
	if eventReq.Location != "" {
		updateMap["location"] = &eventReq.Location
	}
	if eventReq.HostedBy != "" {
		updateMap["hosted_by"] = &eventReq.HostedBy
	}
	if eventReq.Image != "" {
		updateMap["url"] = &eventReq.Image
	}

	err = uc.eventData.UpdateDataByID(updateMap, id, userId)
	return err
}

func (uc *eventUseCase) GetEventByUserID(id_user, limit, page int) (response []events.Core, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.eventData.SelectDataByUserID(id_user, limit, offset)
	total = total/int64(limit) + 1
	return resp, total, errData
}
