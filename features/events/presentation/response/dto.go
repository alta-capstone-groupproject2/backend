package response

import (
	"lami/app/features/events"
	"time"
)

type Event struct {
	ID       int       `json:"eventID" form:"eventID"`
	Image    string    `json:"image" form:"image"`
	Document string    `json:"document" form:"document"`
	Name     string    `json:"eventName" form:"eventName"`
	HostedBy string    `json:"hostedBy" form:"hostedBy"`
	Date     time.Time `json:"date" form:"date"`
	City     string    `json:"city" form:"city"`
	Location string    `json:"location" form:"location"`
	Detail   string    `json:"details" form:"details"`
	Price    int       `json:"price" form:"price"`
}

type EventByID struct {
	ID          int           `json:"eventID" form:"eventID"`
	Image       string        `json:"image" form:"image"`
	Document    string        `json:"document" form:"document"`
	Name        string        `json:"eventName" form:"eventName"`
	HostedBy    string        `json:"hostedby" form:"hostedby"`
	Phone       string        `json:"phone" form:"phone"`
	Date        time.Time     `json:"date" form:"date"`
	City        string        `json:"city" form:"city"`
	Location    string        `json:"location" form:"location"`
	Detail      string        `json:"details" form:"details"`
	Price       int           `json:"price" form:"price"`
	Participant []Participant `json:"participant" form:"participant"`
}

type Participant struct {
	ID    int    `json:"participantID" form:"participantID"`
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}

func FromCore(data events.Core) Event {
	return Event{
		ID:       data.ID,
		Image:    data.Image,
		Document: data.Document,
		Name:     data.Name,
		Detail:   data.Detail,
		HostedBy: data.HostedBy,
		Date:     data.Date,
		City:     data.City,
		Location: data.Location,
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
		Date:        data.Date,
		City:        data.City,
		Location:    data.Location,
		Price:       data.Price,
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

func FromParticipantCoreList(data []events.Participant) []Participant {
	result := []Participant{}
	for key := range data {
		result = append(result, FromParticipantCore(data[key]))
	}
	return result
}
