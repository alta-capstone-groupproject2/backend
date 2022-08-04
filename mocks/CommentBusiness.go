package mocks

import (
	"lami/app/features/comments"

	mock "github.com/stretchr/testify/mock"
)

type CommentUseCase struct {
	mock.Mock
}

func (_m *CommentUseCase) AddComment(dataEvent comments.Core) (err error) {
	ret := _m.Called(dataEvent)

	var r0 error
	if rf, ok := ret.Get(0).(func(comments.Core) error); ok {
		r0 = rf(dataEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *CommentUseCase) GetCommentByIdEvent(limit, page, eventId int) (res []comments.Core, total int64, err error) {
	ret := _m.Called(limit, page, eventId)

	var r0 []comments.Core
	if rf, ok := ret.Get(0).(func(int, int, int) []comments.Core); ok {
		r0 = rf(limit, page, eventId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Core)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int, int) int64); ok {
		r1 = rf(limit, page, eventId)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
		r2 = rf(limit, page, eventId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewCommentBusiness interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserBusiness creates a new instance of UserBusiness. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentBusiness(t mockConstructorTestingTNewCommentBusiness) *CommentUseCase {
	mock := &CommentUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
