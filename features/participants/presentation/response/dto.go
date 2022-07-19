package response

import (
	"time"

	"lami/app/features/participants"
)

type Participant struct {
	ID            int       `json:"participantID" form:"participantID"`
	Name          string    `json:"name" form:"name"`
	Detail        string    `json:"details" form:"details"`
	Image         string    `json:"image" form:"image"`
	Date          time.Time `json:"date" form:"date"`
	HostedBy      string    `json:"hostedby" form:"hostedby"`
	City          string    `json:"city" form:"city"`
	Location      string    `json:"location" form:"location"`
	GrossAmount   int64     `json:"gross_amount"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
}

type Payment struct {
	OrderID           string    `json:"order_id"`
	TransactionID     string    `json:"transaction_id"`
	PaymentMethod     string    `json:"payment_method"`
	BillNumber        string    `json:"bill_number"`
	Bank              string    `json:"bank"`
	GrossAmount       int64     `json:"gross_amount"`
	TransactionTime   time.Time `json:"transaction_time"`
	TransactionExpire time.Time `json:"transaction_expire"`
}

func FromCore(core participants.Core) Participant {
	return Participant{
		ID:            core.ID,
		Name:          core.Event.Name,
		Detail:        core.Event.Detail,
		Image:         core.Event.Image,
		Date:          core.Event.Date,
		HostedBy:      core.Event.HostedBy,
		City:          core.Event.City,
		Location:      core.Event.Location,
		GrossAmount:   core.GrossAmount,
		PaymentMethod: core.PaymentMethod,
		TransactionID: core.TransactionID,
		Status:        core.Status,
	}
}

func FromCoreList(data []participants.Core) []Participant {
	result := []Participant{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
