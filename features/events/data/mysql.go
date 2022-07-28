package data

import (
	"errors"
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
	result := repo.db.Order("id desc").Where("city LIKE ? and name LIKE ? and status = ?", "%"+city+"%", "%"+name+"%", config.Approved).Limit(limit).Offset(offset).Find(&dataEvent).Count(&count)
	if result.Error != nil {
		return []events.Core{}, 0, result.Error
	}
	return ToCoreList(dataEvent), count, result.Error
}

func (repo *mysqlEventRepository) SelectDataByID(id int) (response events.Core, err error) {
	dataEvent := Event{}
	result := repo.db.Where("status = ?", config.Approved).Find(&dataEvent, id)
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
	result := repo.db.Where("user_id = ?", userId).Delete(&dataEvent, id)
	if result.RowsAffected == 0 {
		return errors.New("no rows affected")
	}
	return result.Error
}

func (repo *mysqlEventRepository) CheckValidateUserID(id int) (respon int, err error) {
	var data Event
	result := repo.db.First(&data, id)
	if result.Error != nil {
		return 0, result.Error
	}
	return data.UserID, nil
}

func (repo *mysqlEventRepository) UpdateDataByID(status string, id, userID int) error {
	model := Event{}
	model.ID = uint(id)

	result := repo.db.Model(&model).Where("user_id = ?", userID).Update("status", status)
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
	result := repo.db.Order("updated_at desc").Limit(limit).Offset(offset).Preload("User").Find(&dataSubmit).Count(&count)
	if result.Error != nil {
		return []events.Submission{}, 0, result.Error
	}
	return ToCoreSubmissionList(dataSubmit), count, result.Error
}

func (repo *mysqlEventRepository) SelectDataSubmissionByID(id int) (data events.Core, err error) {
	var dataSubmit Event
	result := repo.db.Where("id = ?", id).Find(&dataSubmit)
	if result.Error != nil {
		return events.Core{}, result.Error
	}
	return dataSubmit.toCore(), result.Error
}

func (repo *mysqlEventRepository) SelectAttendeeData(id_event int) (response []events.AttendeesData, err error) {
	var dataParticipant []Participant

	result := repo.db.Preload("User").Find(&dataParticipant, "event_id = ?", id_event)
	if result.Error != nil {
		return []events.AttendeesData{}, result.Error
	}

	return ToAttendeeCoreList(dataParticipant), result.Error
}
