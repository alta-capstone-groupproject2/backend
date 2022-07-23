package data

import (
	"lami/app/features/events"
	"lami/app/features/users/data"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Image     string
	Document  string
	Name      string
	HostedBy  string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	City      string
	Location  string
	Detail    string
	Price     int
	Status    string
	UserID    int
	User      data.User
}

type Participant struct {
	ID      int
	UserID  int
	EventID int
	Name    string
	Image   string
	User    data.User
	Event   Event
}

//DTO

func (data *Event) toCore() events.Core {
	return events.Core{
		ID:        int(data.ID),
		Document:  data.Document,
		Image:     data.Image,
		Name:      data.Name,
		HostedBy:  data.HostedBy,
		Phone:     data.Phone,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		City:      data.City,
		Location:  data.Location,
		Detail:    data.Detail,
		Price:     data.Price,
		Status:    data.Status,
		UserID:    data.UserID,
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
		Image:     core.Image,
		Document:  core.Document,
		Name:      core.Name,
		HostedBy:  core.HostedBy,
		Phone:     core.Phone,
		StartDate: core.StartDate,
		EndDate:   core.EndDate,
		City:      core.City,
		Location:  core.Location,
		Detail:    core.Detail,
		Price:     core.Price,
		Status:    core.Status,
		UserID:    core.UserID,
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

func (data *Event) toSubmissionCore() events.Submission {
	return events.Submission{
		ID:        int(data.ID),
		Name:      data.Name,
		UserName:  data.User.Name,
		City:      data.City,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Status:    data.Status,
	}
}

func ToCoreSubmissionList(data []Event) []events.Submission {
	result := []events.Submission{}
	for key := range data {
		result = append(result, data[key].toSubmissionCore())
	}
	return result
}

func (data *Participant) toAttendeeCore() events.AttendeesData {
	return events.AttendeesData{
		Name:    data.User.Name,
		Email:   data.User.Email,
		City:    data.User.City,
		Present: "",
	}
}

func ToAttendeeCoreList(data []Participant) []events.AttendeesData {
	result := []events.AttendeesData{}
	for key := range data {
		result = append(result, data[key].toAttendeeCore())
	}
	return result
}
