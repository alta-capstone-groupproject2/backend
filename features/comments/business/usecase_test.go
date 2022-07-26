package business

import (
	"lami/app/features/comments"
	mockComment "lami/app/mocks/comments"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock data success case
// type mockCommentDataSuccess struct{}

// func (mock mockCommentDataSuccess) Insert(data comments.Core) (row int, err error) {
// 	return 1, nil
// }

// func (mock mockCommentDataSuccess) GetComment(limit, offset, event_id int) (data []comments.Core, count int64, err error) {
// 	return []comments.Core{
// 		{ID: 1, UserID: 1, EventID: 1, Comment: "Bisa"},
// 	}, 1, nil
// }

// // mock data failed case
// type mockCommentDataFailed struct{}

// func (mock mockCommentDataFailed) Insert(data comments.Core) (row int, err error) {
// 	return 0, fmt.Errorf("failed insert your comment")
// }

// func (mock mockCommentDataFailed) GetComment(limit, offset, event_id int) (data []comments.Core, count int64, err error) {
// 	return nil, 0, fmt.Errorf("No data")
// }

func TestAddComment(t *testing.T) {
	repo := new(mockComment.CommentBusiness)
	insertData := comments.Core{
		EventID: 1,
		UserID:  1,
		Comment: "Apakah masih bisa ?",
	}

	t.Run("Test Add Comment Success", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()
		srv := NewCommentBusiness(repo)

		res, err := srv.AddComment(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	// t.Run("Test Add Comment Failed", func(t *testing.T) {
	// 	commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
	// 	eventId := 1
	// 	userId := 1
	// 	Comment := "apakah masih bisa ?"
	// 	addComment := comments.Core{
	// 		EventID: eventId,
	// 		UserID:  userId,
	// 		Comment: Comment,
	// 	}
	// 	result, err := commentBusiness.AddComment(addComment)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, 0, result)
	// })
}

// func TestGetCommentByIdEvent(t *testing.T) {
// 	t.Run("Test Get Comment By Id Event Success", func(t *testing.T) {
// 		commentBusiness := NewCommentBusiness(mockCommentDataSuccess{})
// 		eventId := 1
// 		limit := 5
// 		offset := 0
// 		result, count, err := commentBusiness.GetCommentByIdEvent(limit, offset, eventId)
// 		assert.Nil(t, err)
// 		assert.Equal(t, int64(1), count)
// 		assert.Equal(t, 1, result[0].EventID)
// 	})
// 	t.Run("Test Get Comment By Id Event Failed", func(t *testing.T) {
// 		commentBusiness := NewCommentBusiness(mockCommentDataFailed{})
// 		eventId := 2
// 		limit := 5
// 		offset := 0
// 		_, count, err := commentBusiness.GetCommentByIdEvent(limit, offset, eventId)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, int64(0), count)
// 	})
// }
