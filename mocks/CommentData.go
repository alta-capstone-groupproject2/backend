package mocks

import (
	"lami/app/features/comments"

	"github.com/stretchr/testify/mock"
)

type CommentData struct {
	mock.Mock
}

func (_m *CommentData) Insert(dataEvent comments.Core) (err error) {
	ret := _m.Called(dataEvent)

	var r1 error
	if rf, ok := ret.Get(0).(func(comments.Core) error); ok {
		r1 = rf(dataEvent)
	} else {
		r1 = ret.Error(0)
	}

	return r1
}

func (_m *CommentData) GetComment(limit, page, eventID int) (res []comments.Core, total int64, err error) {
	ret := _m.Called(limit, page, eventID)

	var r0 []comments.Core
	if rf, ok := ret.Get(0).(func(int, int, int) []comments.Core); ok {
		r0 = rf(limit, page, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Core)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(int, int, int) int64); ok {
		r1 = rf(limit, page, eventID)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int) error); ok {
		r2 = rf(limit, page, eventID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewCommentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthData creates a new instance of AuthData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentData(t mockConstructorTestingTNewCommentData) *CommentData {
	mock := &CommentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
