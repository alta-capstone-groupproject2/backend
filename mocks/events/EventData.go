package events

import (
	"lami/app/features/events"

	mock "github.com/stretchr/testify/mock"
)

type EventData struct {
	mock.Mock
}

func (_m *EventData) SelectData(limit, offset int, city, name string) (data []events.Core, total int64, err error) {
	ret := _m.Called(limit, offset, city, name)

	var r0 []events.Core
	if rf, ok := ret.Get(0).(func(int, int, string, string) []events.Core); ok {
		r0 = rf(limit, offset, city, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Core)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int, string, string) int64); ok {
		r1 = rf(limit, offset, city, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(int64)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, string, string) error); ok {
		r2 = rf(limit, offset, city, name)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *EventData) SelectDataByID(eventID int) (dataEvent events.Core, err error) {
	ret := _m.Called(eventID)

	var r0 events.Core
	if rf, ok := ret.Get(0).(func(int) events.Core); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Get(0).(events.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *EventData) InsertData(dataReq events.Core) (err error) {
	ret := _m.Called(dataReq)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.Core) error); ok {
		r0 = rf(dataReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventData) DeleteDataByID(eventID, userID int) (err error) {
	ret := _m.Called(eventID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(eventID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventData) UpdateDataByID(status string, eventID, userID int) (err error) {
	ret := _m.Called(status, eventID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int, int) error); ok {
		r0 = rf(status, eventID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventData) SelectDataByUserID(userID, limit, offset int) (dataEvent []events.Core, total int64, err error) {
	ret := _m.Called(userID, limit, offset)

	var r0 []events.Core
	if rf, ok := ret.Get(0).(func(int, int, int) []events.Core); ok {
		r0 = rf(userID, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Core)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int, int) int64); ok {
		r1 = rf(userID, limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(int64)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
		r2 = rf(userID, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *EventData) SelectParticipantData(eventID int) (dataEvent []events.Participant, err error) {
	ret := _m.Called(eventID)

	var r0 []events.Participant
	if rf, ok := ret.Get(0).(func(int) []events.Participant); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Get(0).([]events.Participant)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *EventData) CheckValidateUserID(eventID int) (userID int, err error) {
	ret := _m.Called(eventID)

	var r0 events.Core
	if rf, ok := ret.Get(0).(func(int) events.Core); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Get(0).(events.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0.UserID, r1
}

func (_m *EventData) SelectDataSubmission(limit, offset int) (dataEvent []events.Submission, total int64, err error) {
	ret := _m.Called(limit, offset)

	var r0 []events.Submission
	if rf, ok := ret.Get(0).(func(int, int) []events.Submission); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Submission)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int) int64); ok {
		r1 = rf(limit, offset)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(int64)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

func (_m *EventData) SelectDataSubmissionByID(eventID int) (dataEvent events.Core, err error) {
	ret := _m.Called(eventID)

	var r0 events.Core
	if rf, ok := ret.Get(0).(func(int) events.Core); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Get(0).(events.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *EventData) SelectAttendeeData(eventID int) (dataEvent []events.AttendeesData, err error) {
	ret := _m.Called(eventID)

	var r0 []events.AttendeesData
	if rf, ok := ret.Get(0).(func(int) []events.AttendeesData); ok {
		r0 = rf(eventID)
	} else {
		r0 = ret.Get(0).([]events.AttendeesData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
