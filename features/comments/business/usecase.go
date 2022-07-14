package business

import (
	"lami/app/features/comments"
)

type commentUseCase struct {
	commentData comments.Data
}

func NewCommentBusiness(cmnData comments.Data) comments.Business {
	return &commentUseCase{
		commentData: cmnData,
	}
}

func (uc *commentUseCase) AddComment(data comments.Core) (row int, err error) {
	row, err = uc.commentData.Insert(data)
	return row, err
}

func (uc *commentUseCase) GetCommentByIdEvent(limit, offset, eventId int) (response []comments.Core, total int64, err error) {
	page := limit * (offset - 1)
	response, total, err = uc.commentData.GetComment(limit, page, eventId)
	if err != nil {
		return []comments.Core{}, 0, err
	}
	total = total/int64(limit) + 1
	return response, total, err
}
