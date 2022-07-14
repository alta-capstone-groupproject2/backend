package data

import (
	"lami/app/features/auth"

	// "project3/eventapp/features/auth/data"

	"gorm.io/gorm"
)

type mysqlAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(conn *gorm.DB) auth.Data {
	return &mysqlAuthRepository{
		db: conn,
	}
}

func (repo *mysqlAuthRepository) FindUser(email string) (response auth.Core, err error) {
	dataUser := User{}
	result := repo.db.Where("email = ?", email).Find(&dataUser)
	if result.Error != nil {
		return auth.Core{}, result.Error
	}

	return dataUser.toCore(), nil
}
