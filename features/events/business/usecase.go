package business

import (
	"errors"
	"lami/app/features/events"
	"lami/app/helper"
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
	if eventRequest.Name == "" || eventRequest.Image == "" || eventRequest.Document == "" || eventRequest.Detail == "" || eventRequest.City == "" || eventRequest.Location == "" || eventRequest.HostedBy == "" {
		return errors.New("all data must be filled")
	}

	err := uc.eventData.InsertData(eventRequest)
	return err
}

func (uc *eventUseCase) DeleteEventByID(id int, userId int) (err error) {
	err = uc.eventData.DeleteDataByID(id, userId)
	return err
}

func (uc *eventUseCase) UpdateEventByID(status string, id int) (err error) {
	checkUserID, errCheck := uc.eventData.CheckUserID(id)
	if errCheck != nil {
		return err
	}
	userID := checkUserID
	err = uc.eventData.UpdateDataByID(status, id, userID)
	return err
}

func (uc *eventUseCase) GetEventByUserID(userID, limit, page int) (response []events.Core, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.eventData.SelectDataByUserID(userID, limit, offset)
	total = total/int64(limit) + 1
	return resp, total, errData
}

func (uc *eventUseCase) GetEventSubmission(limit, page int) (data []events.Submission, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.eventData.SelectDataSubmission(limit, offset)
	total = total/int64(limit) + 1
	return resp, total, errData
}

func (uc *eventUseCase) GetEventSubmissionByID(id int) (data events.Core, err error) {
	result, err := uc.eventData.SelectDataSubmissionByID(id)
	if err != nil {
		return events.Core{}, err
	}
	return result, err
}

func (uc *eventUseCase) GetEventAttendee(id int, userID int) (urlPDF string, err error) {
	response, err := uc.eventData.SelectDataByID(id)
	if err != nil {
		return "", err
	}
	checkUserID, errCheck := uc.eventData.CheckUserID(id)
	if errCheck != nil {
		return "", err
	}
	if checkUserID != userID {
		return "", errors.New("not found")
	}
	responseAttendee, errAttendee := uc.eventData.SelectAttendeeData(response.ID)
	if errAttendee != nil {
		return "", errAttendee
	}
	response.AttendeesData = responseAttendee
	urlPDF = helper.ExportPDF(response)
	return urlPDF, err
}
