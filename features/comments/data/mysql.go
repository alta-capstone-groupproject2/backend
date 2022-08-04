package data

import (
	"lami/app/features/comments"

	"gorm.io/gorm"
)

type mysqlCommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(conn *gorm.DB) comments.Data {
	return &mysqlCommentRepository{
		db: conn,
	}
}

func (repo *mysqlCommentRepository) Insert(dataReq comments.Core) (err error) {
	Comment := fromCore(dataReq)
	result := repo.db.Create(&Comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *mysqlCommentRepository) GetComment(limit, offset, eventID int) (response []comments.Core, total int64, err error) {
	var dataComment []Comment
	var count int64
	result := repo.db.Order("id DESC").Where("event_id = ?", eventID).Preload("User").Limit(limit).Offset(offset).Find(&dataComment).Count(&count)

	if result.Error != nil {
		return []comments.Core{}, int64(0), result.Error
	}
	return ToCoreList(dataComment), count, result.Error
}
