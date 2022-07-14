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

func (repo *mysqlCommentRepository) Insert(data comments.Core) (row int, err error) {
	commentData := fromCore(data)
	result := repo.db.Create(&commentData)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(result.RowsAffected), result.Error
}

func (repo *mysqlCommentRepository) GetComment(limit, offset, eventId int) (response []comments.Core, total int64, err error) {
	var dataComment []Comment
	var count int64
	result := repo.db.Order("id DESC").Where("event_id = ?", eventId).Preload("User").Limit(limit).Offset(offset).Find(&dataComment).Count(&count)

	if result.Error != nil {
		return []comments.Core{}, int64(0), result.Error
	}
	return ToCoreList(dataComment), count, result.Error
}
