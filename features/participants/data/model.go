package data

import (
	_event "lami/app/features/events/data"
	"lami/app/features/participants"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID  int
	EventID int
	Event   _event.Event
}

func fromCore(core participants.Core) Participant {
	return Participant{
		UserID:  core.UserID,
		EventID: core.EventID,
	}
}

func (data *Participant) toCore() participants.Core {
	return participants.Core{
		ID: int(data.ID),
		Event: participants.Event{
			Image:    data.Event.Image,
			Name:     data.Event.Name,
			HostedBy: data.Event.HostedBy,
			Date:     data.Event.Date,
			City:     data.Event.City,
			Location: data.Event.Location,
			Detail:   data.Event.Detail,
		},
	}
}

func ToCoreList(data []Participant) []participants.Core {
	result := []participants.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
