package events

import (
	"lami/app/features/events"

	"github.com/stretchr/testify/mock"
)

type EventBusiness struct {
	mock.Mock
}

func (_m *EventBusiness) GetAllEvent(limit, offset int, city, name string) (dataEvent []events.Core, total int64, err error) {
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

func (_m *EventBusiness) GetEventByID(eventID int) (dataEvent events.Core, err error) {
	ret := _m.Called(eventID)

	var r0 events.Core
	if rf, ok := ret.Get(0).(func(int) events.Core); ok {
		r0 = rf(eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(events.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

func (_m *EventBusiness) InsertEvent(dataReq events.Core) (err error) {
	ret := _m.Called(dataReq)

	var r0 error
	if rf, ok := ret.Get(0).(func(events.Core) error); ok {
		r0 = rf(dataReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventBusiness) DeleteEventByID(eventID, userID int) (err error) {
	ret := _m.Called(eventID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(eventID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventBusiness) UpdateEventByID(status string, eventID int) (err error) {
	ret := _m.Called(status, eventID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(status, eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *EventBusiness) GetEventByUserID(userID, limit, offset int) (dataEvent []events.Core, total int64, err error) {
	ret := _m.Called(limit, offset)

	var r0 []events.Core
	if rf, ok := ret.Get(0).(func(int, int) []events.Core); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Core)
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

func (_m *EventBusiness) GetEventSubmission(limit, offset int) (dataEvent []events.Core, total int64, err error) {
	ret := _m.Called(limit, offset)

	var r0 []events.Core
	if rf, ok := ret.Get(0).(func(int, int) []events.Core); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]events.Core)
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

func (_m *EventBusiness) GetEventSubmissionByID(eventID int) (dataEvent events.Core, err error) {
	ret := _m.Called(eventID)

	var r0 events.Core
	if rf, ok := ret.Get(0).(func(int) events.Core); ok {
		r0 = rf(eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(events.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(int) error); ok {
		r1 = rf(eventID)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

func (_m *EventBusiness) GetEventAttendee(eventID, userID int) (urlPDF string, err error) {
	ret := _m.Called(eventID, userID)

	var r0 string
	if rf, ok := ret.Get(0).(func(int, int) string); ok {
		r0 = rf(eventID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r1 = rf(eventID, userID)
	} else {
		r1 = ret.Error(2)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventBussiness interface {
	mock.TestingT
	Cleanup(func())
}

// NewCultureBussiness creates a new instance of CultureBussiness. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventBussiness(t mockConstructorTestingTNewEventBussiness) *EventBusiness {
	mock := &EventBusiness{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
