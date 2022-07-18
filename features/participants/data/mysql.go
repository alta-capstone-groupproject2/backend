package data

import (
	"errors"
	"lami/app/config"
	_eventData "lami/app/features/events/data"
	"lami/app/features/participants"

	"gorm.io/gorm"
)

type mysqlParticipantRepository struct {
	db *gorm.DB
}

func NewParticipantRepository(conn *gorm.DB) participants.Data {
	return &mysqlParticipantRepository{
		db: conn,
	}
}

func (repo *mysqlParticipantRepository) SelectDataByID(id int) (response _eventData.Event, err error) {
	dataEvent := _eventData.Event{}
	result := repo.db.Where("status = ?", config.Approved).Find(&dataEvent, id)
	if result.Error != nil {
		return _eventData.Event{}, result.Error
	}

	return dataEvent, err
}

// DeleteData implements participants.Data
func (repo *mysqlParticipantRepository) DeleteData(param, userID int) error {
	dataparticipant := Participant{}
	result := repo.db.Where("event_id = ? AND user_id = ?", param, userID).Delete(&dataparticipant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// SelectDataEvent implements participants.Data
func (repo *mysqlParticipantRepository) SelectDataEvent(idUser int) (data []participants.Core, err error) {
	dataParticipant := []Participant{}

	result := repo.db.Preload("Event").Where("user_id = ?", idUser).Find(&dataParticipant)
	if result.Error != nil {
		return []participants.Core{}, result.Error
	}

	return ToCoreList(dataParticipant), result.Error
}

// Add implements participants.Data
func (repo *mysqlParticipantRepository) AddData(ParticipantData participants.Core) error {
	Model := fromCore(ParticipantData)
	result := repo.db.Create(&Model)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed insert join")
	}
	return nil
}
