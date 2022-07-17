package request

import (
	"lami/app/features/events"
	"time"
)

type Event struct {
	Image    string `json:"image" form:"image"`
	Document string `json:"document" form:"document"`
	Name     string `json:"name" form:"name"`
	HostedBy string `json:"hostedBy" form:"hostedBy"`
	Phone    string `json:"phone" form:"phone"`
	Date     string `json:"date" form:"date"`
	City     string `json:"city" form:"city"`
	Location string `json:"location" form:"location"`
	Detail   string `json:"details" form:"details"`
	Price    int    `json:"price" form:"price"`
	DateTime time.Time
}

type UpdateEvent struct {
	Status string `json:"status" form:"status"`
}

func ToCore(eventReq Event) events.Core {
	eventCore := events.Core{
		Image:    eventReq.Image,
		Document: eventReq.Document,
		Name:     eventReq.Name,
		HostedBy: eventReq.HostedBy,
		Phone:    eventReq.Phone,
		Date:     eventReq.DateTime,
		City:     eventReq.City,
		Location: eventReq.Location,
		Detail:   eventReq.Detail,
		Price:    eventReq.Price,
	}
	return eventCore
}
