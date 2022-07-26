package comments

import (
	"lami/app/features/comments"

	mock "github.com/stretchr/testify/mock"
)

type CommentBusiness struct {
	mock.Mock
}

func (_m *CommentBusiness) Insert(data comments.Core) (row int, err error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(comments.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comments.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CommentBusiness) GetComment(limit, offset, eventId int) (res []comments.Core, total int64, err error) {
	ret := _m.Called(limit, offset, eventId)
	page := limit * (offset - 1)

	var r0 []comments.Core
	var r1 int64
	if rf, ok := ret.Get(0).(func(int, int, int) []comments.Core); ok {
		r0 = rf(limit, page, eventId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Core)
		}
	}

	var r2 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(1)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewCommentBusiness interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserBusiness creates a new instance of UserBusiness. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentBusiness(t mockConstructorTestingTNewCommentBusiness) *CommentBusiness {
	mock := &CommentBusiness{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
