package business

import (
	"errors"
	"lami/app/features/comments"
	mockComment "lami/app/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddComment(t *testing.T) {
	repo := new(mockComment.CommentData)
	dataReq := comments.Core{
		EventID: 1,
		UserID:  1,
		Comment: "Apakah masih bisa ?",
	}

	t.Run("Test Add Comment Success", func(t *testing.T) {
		repo.On("Insert", dataReq).Return(nil).Once()
		srv := NewCommentBusiness(repo)

		err := srv.AddComment(dataReq)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Test Add Comment Failed", func(t *testing.T) {
		dataReq.EventID = 0
		repo.On("Insert", dataReq).Return(errors.New("failed add comment")).Once()
		srv := NewCommentBusiness(repo)

		err := srv.AddComment(dataReq)
		assert.NotNil(t, err)
	})
}

func TestGetComment(t *testing.T) {
	repo := new(mockComment.CommentData)
	returnData := []comments.Core{
		{
			ID:      1,
			EventID: 1,
			UserID:  1,
			Comment: "Test 1",
		},
		{
			ID:      2,
			EventID: 1,
			UserID:  2,
			Comment: "Test 2",
		},
	}

	t.Run("Test Get Comment Success", func(t *testing.T) {
		var total int64 = 1
		repo.On("GetComment", 5, 0, 1).Return(returnData, total, nil).Once()
		srv := NewCommentBusiness(repo)

		result, sum, err := srv.GetCommentByIdEvent(5, 1, 1)
		totalpage := sum/int64(5) + 1
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].ID, result[0].ID)
		assert.Equal(t, total, totalpage)
		repo.AssertExpectations(t)
	})

	t.Run("Test Get Comment Failed", func(t *testing.T) {
		var total int64 = 1
		returnDataFailed := []comments.Core{}
		repo.On("GetComment", 5, 0, 3).Return(returnDataFailed, total, errors.New("no data response")).Once()
		srv := NewCommentBusiness(repo)

		result, sum, err := srv.GetCommentByIdEvent(5, 1, 3)
		totalpage := sum/int64(5) + 1
		assert.NotNil(t, err)
		assert.Equal(t, returnDataFailed, result)
		assert.Equal(t, total, totalpage)
	})
}
