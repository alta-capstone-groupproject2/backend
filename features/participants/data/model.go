package data

import (
	_event "lami/app/features/events/data"
	"lami/app/features/participants"

	"gorm.io/gorm"
)

type Participant struct {
	gorm.Model
	UserID        int
	EventID       int
	OrderID       string
	GrossAmount   int64
	PaymentMethod string
	TransactionID string
	Status        string
	Event         _event.Event
}

func fromCore(core participants.Core) Participant {
	return Participant{
		UserID:        core.UserID,
		EventID:       core.EventID,
		OrderID:       core.OrderID,
		GrossAmount:   int64(core.Event.Price),
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

func (data *Participant) toCoreMidtrans() participants.Core {
	return participants.Core{
		ID:            int(data.ID),
		UserID:        data.UserID,
		EventID:       data.EventID,
		OrderID:       data.OrderID,
		GrossAmount:   int64(data.Event.Price),
		PaymentMethod: data.PaymentMethod,
		TransactionID: data.TransactionID,
		Status:        data.Status,
	}
}
