package data

import (
	"lami/app/features/events"
	"lami/app/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Image    string
	Document string
	Name     string
	HostedBy string
	Phone    string
	Date     time.Time
	City     string
	Location string
	Detail   string
	Price    int
	Status   string
	UserID   int
	User     data.User
}

type Participant struct {
	ID      int
	UserID  int
	EventID int
	Name    string
	Image   string
	User    data.User
}

//DTO

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:       int(data.ID),
		Document: data.Document,
		Image:    data.Image,
		Name:     data.Name,
		HostedBy: data.HostedBy,
		Phone:    data.Phone,
		Date:     data.Date,
		City:     data.City,
		Location: data.Location,
		Detail:   data.Detail,
		Price:    data.Price,
		Status:   data.Status,
		IDUser:   data.UserID,
	}
}

func ToCoreList(data []Event) []events.Core {
	result := []events.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func fromCore(core events.Core) Event {
	return Event{
		Image:    core.Image,
		Document: core.Document,
		Name:     core.Name,
		HostedBy: core.HostedBy,
		Phone:    core.Phone,
		Date:     core.Date,
		City:     core.City,
		Location: core.Location,
		Detail:   core.Detail,
		Price:    core.Price,
		Status:   core.Status,
		UserID:   core.IDUser,
	}
}

func toCore(data Event) events.Core {
	return data.toCore()
}

func (data *Participant) toParticipantCore() events.Participant {
	return events.Participant{
		ID:    data.ID,
		Name:  data.User.Name,
		Image: data.User.Image,
	}
}

func ToParticipantCoreList(data []Participant) []events.Participant {
	result := []events.Participant{}
	for key := range data {
		result = append(result, data[key].toParticipantCore())
	}
	return result
}
