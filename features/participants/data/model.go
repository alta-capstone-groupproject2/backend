package data

import (
	_event "lami/app/features/events/data"
	"lami/app/features/participants"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID        int
	EventID       int
	OrderID       string
	GrossAmount   string
	PaymentMethod string
	TransactionID string
	Status        string
	Event         _event.Event
}

type Payment struct {
	OrderID           string
	TransactionID     string
	PaymentMethod     string
	BillNumber        string
	Bank              string
	GrossAmount       int64
	TransactionTime   time.Time
	TransactionExpire time.Time
}

func fromCore(core participants.Core) Participant {
	return Participant{
		UserID:        core.UserID,
		EventID:       core.EventID,
		OrderID:       core.OrderID,
		GrossAmount:   strconv.Itoa(core.Event.Price),
		PaymentMethod: core.PaymentMethod,
		TransactionID: core.TransactionID,
		Status:        core.Status,
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
			Price:    data.Event.Price,
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
