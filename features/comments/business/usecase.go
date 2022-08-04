package business

import (
	"lami/app/features/comments"
)

type CommentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(cmnData comments.Data) comments.Business {
	return &CommentUseCase{
		commentData: cmnData,
	}
}

func (uc *CommentUseCase) AddComment(dataReq comments.Core) (err error) {
	result := uc.commentData.Insert(dataReq)
	if result != nil {
		return result
	}
	return nil
}

func (uc *CommentUseCase) GetCommentByIdEvent(limit, offset, eventID int) (response []comments.Core, total int64, err error) {
	page := limit * (offset - 1)
	response, total, err = uc.commentData.GetComment(limit, page, eventID)
	if err != nil {
		return []comments.Core{}, 0, err
	}
	total = total/int64(limit) + 1
	return response, total, err
}
