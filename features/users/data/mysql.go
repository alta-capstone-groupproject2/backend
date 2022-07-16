package data

import (
	"encoding/json"
	"errors"

	"lami/app/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) users.Data {
	return &mysqlUserRepository{
		db: conn,
	}
}

func (repo *mysqlUserRepository) SelectDataById(id int) (response users.Core, err error) {
	datauser := User{}
	result := repo.db.Preload("Role").Find(&datauser, id)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	if result.RowsAffected == 0 {
		return users.Core{}, errors.New("user not found")
	}
	return toCore(datauser), nil
}

func (repo *mysqlUserRepository) InsertData(userData users.Core) (err error) {
	userModel := fromCore(userData)
	var newError map[string]interface{}

	result := repo.db.Create(&userModel)

	if result.Error != nil {
		errByte, _ := json.Marshal(result.Error)
		json.Unmarshal((errByte), &newError)

		if newError["Number"] == float64(1062) {
			return errors.New("email already used")
		}

		return errors.New("failed insert data")
	}
	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}
	return nil
}

func (repo *mysqlUserRepository) DeleteData(id int) (err error) {
	datauser := User{}
	result := repo.db.Delete(&datauser, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete data")
	}
	return nil
}

func (repo *mysqlUserRepository) UpdateData(dataReq map[string]interface{}, id int) (err error) {
	model := User{}
	model.ID = uint(id)
	result := repo.db.Model(model).Updates(dataReq)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update data")
	}
	return nil
}

func (repo *mysqlUserRepository) InsertStoreData(dataReq users.Core, id int) error {
	model := User{}
	model.ID = uint(id)
	dataStore := fromCore(dataReq)
	result := repo.db.Model(model).Updates(dataStore)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update store data")
	}
	return nil
}

func (repo *mysqlUserRepository) UpdateAccountRole(status string, id int) error {
	model := User{}
	model.ID = uint(id)
	result := repo.db.Model(model).Updates(map[string]interface{}{"store_status": status, "role_id": 3})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update data")
	}
	return nil
}

func (repo *mysqlUserRepository) SelectData(limit, offset int) (response []users.Core, total int64, err error) {
	var dataUser []User
	var count int64
	result := repo.db.Preload("Role").Where("store_status is not null").Find(&dataUser).Count(&count)
	if result.Error != nil {
		return []users.Core{}, 0, result.Error
	}
	return ToCoreList(dataUser), count, result.Error
}
