package response

import (
	"lami/app/features/events"
	"time"
)

type Event struct {
	ID        int       `json:"eventID" form:"eventID"`
	Image     string    `json:"image" form:"image"`
	Document  string    `json:"document" form:"document"`
	Name      string    `json:"eventName" form:"eventName"`
	HostedBy  string    `json:"hostedBy" form:"hostedBy"`
	StartDate time.Time `json:"startDate" form:"startDate"`
	EndDate   time.Time `json:"endDate" form:"endDate"`
	City      string    `json:"city" form:"city"`
	Location  string    `json:"location" form:"location"`
	Detail    string    `json:"details" form:"details"`
	Price     int       `json:"price" form:"price"`
	Status    string    `json:"status"`
}

type EventByID struct {
	ID          int           `json:"eventID" form:"eventID"`
	Image       string        `json:"image" form:"image"`
	Document    string        `json:"document" form:"document"`
	Name        string        `json:"eventName" form:"eventName"`
	HostedBy    string        `json:"hostedBy" form:"hostedBy"`
	Phone       string        `json:"phone" form:"phone"`
	StartDate   time.Time     `json:"startDate" form:"startDate"`
	EndDate     time.Time     `json:"endDate" form:"endDate"`
	City        string        `json:"city" form:"city"`
	Location    string        `json:"location" form:"location"`
	Detail      string        `json:"details" form:"details"`
	Price       int           `json:"price" form:"price"`
	Status      string        `json:"status" form:"status"`
	Participant []Participant `json:"participant" form:"participant"`
}

type Participant struct {
	ID    int    `json:"participantID" form:"participantID"`
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}

type Submission struct {
	ID        int       `json:"eventID" form:"eventID"`
	Name      string    `json:"nameEvent" form:"nameEvent"`
	UserName  string    `json:"username" form:"username"`
	City      string    `json:"city" form:"city"`
	StartDate time.Time `json:"startDate" form:"startDate"`
	EndDate   time.Time `json:"endDate" form:"endDate"`
	Status    string    `json:"status" form:"status"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:        data.ID,
		Image:     data.Image,
		Document:  data.Document,
		Name:      data.Name,
		Detail:    data.Detail,
		HostedBy:  data.HostedBy,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		City:      data.City,
		Location:  data.Location,
		Price:     data.Price,
		Status:    data.Status,
	}
}

func FromCoreList(data []events.Core) []Event {
	result := []Event{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}

func FromCoreByID(data events.Core) EventByID {
	return EventByID{
		ID:          data.ID,
		Image:       data.Image,
		Document:    data.Document,
		Name:        data.Name,
		Detail:      data.Detail,
		HostedBy:    data.HostedBy,
		Phone:       data.Phone,
		StartDate:   data.StartDate,
		EndDate:     data.EndDate,
		City:        data.City,
		Location:    data.Location,
		Price:       data.Price,
		Status:      data.Status,
		Participant: FromParticipantCoreList(data.Participant),
	}
}

func FromParticipantCore(data events.Participant) Participant {
	return Participant{
		ID:    data.ID,
		Name:  data.Name,
		Image: data.Image,
	}
}

func FromSubmissionCore(data events.Submission) Submission {
	return Submission{
		ID:        data.ID,
		Name:      data.Name,
		UserName:  data.UserName,
		City:      data.City,
		StartDate: data.StartDate,
		EndDate:   data.EndDate,
		Status:    data.Status,
	}
}

func FromParticipantCoreList(data []events.Participant) []Participant {
	result := []Participant{}
	for key := range data {
		result = append(result, FromParticipantCore(data[key]))
	}
	return result
}

func FromSubmissionCoreList(data []events.Submission) []Submission {
	result := []Submission{}
	for key := range data {
		result = append(result, FromSubmissionCore(data[key]))
	}
	return result
}
