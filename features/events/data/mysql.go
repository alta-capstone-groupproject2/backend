package data

import (
	"errors"
	"fmt"
	"lami/app/config"
	"lami/app/features/events"

	"gorm.io/gorm"
)

type mysqlEventRepository struct {
	db *gorm.DB
}

func NewEventRepository(conn *gorm.DB) events.Data {
	return &mysqlEventRepository{
		db: conn,
	}
}

func (repo *mysqlEventRepository) SelectData(limit int, offset int, name string, city string) (response []events.Core, totaldata int64, err error) {
	var dataEvent []Event
	var count int64
	result := repo.db.Order("id desc").Where("city LIKE ? and name LIKE ?", "%"+city+"%", "%"+name+"%").Limit(limit).Offset(offset).Find(&dataEvent).Count(&count)
	if result.Error != nil {
		return []events.Core{}, 0, result.Error
	}
	return ToCoreList(dataEvent), count, result.Error
}

func (repo *mysqlEventRepository) SelectDataByID(id int) (response events.Core, err error) {
	dataEvent := Event{}
	result := repo.db.Find(&dataEvent, id)
	if result.Error != nil {

		return events.Core{}, result.Error
	}

	return toCore(dataEvent), err
}

func (repo *mysqlEventRepository) InsertData(EventData events.Core) error {
	EventModel := fromCore(EventData)
	result := repo.db.Create(&EventModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed insert data")
	}
	return nil
}

func (repo *mysqlEventRepository) DeleteDataByID(id int, userId int) error {
	dataEvent := Event{}
	result := repo.db.Where("user_id = ?", dataEvent.UserID).Delete(&dataEvent, id)
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return result.Error
}

func (repo *mysqlEventRepository) UpdateDataByID(status string, id, userId int) error {
	model := Event{}
	model.ID = uint(id)
	fmt.Println(model.UserID)
	result := repo.db.Model(&model).Where("user_id = ?", model.UserID).Update("status", status)
	if result.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	if result != nil {
		return result.Error
	}

	return nil
}

func (repo *mysqlEventRepository) SelectDataByUserID(id_user, limit, offset int) (response []events.Core, total int64, err error) {
	var dataEvent []Event
	var count int64
	result := repo.db.Where("user_id = ?", id_user).Limit(limit).Offset(offset).Find(&dataEvent).Count(&count)
	if result.Error != nil {
		return []events.Core{}, 0, result.Error
	}
	return ToCoreList(dataEvent), count, result.Error
}

func (repo *mysqlEventRepository) SelectParticipantData(id_event int) (response []events.Participant, err error) {
	var dataParticipant []Participant

	result := repo.db.Order("id desc").Preload("User").Find(&dataParticipant, "event_id = ?", id_event)
	if result.Error != nil {
		return []events.Participant{}, result.Error
	}

	return ToParticipantCoreList(dataParticipant), result.Error
}

func (repo *mysqlEventRepository) SelectDataSubmission(limit, offset int) (data []events.Submission, total int64, err error) {
	var dataSubmit []Event
	var count int64
	result := repo.db.Where("status = ?", config.Status).Limit(limit).Offset(offset).Preload("User").Find(&dataSubmit).Count(&count)
	if result.Error != nil {
		return []events.Submission{}, 0, result.Error
	}
	return ToCoreSubmissionList(dataSubmit), count, result.Error
}

// func (repo *mysqlEventRepository) SelectDataSubmissionByID(dataReq events.Core) (err error) {

// }
