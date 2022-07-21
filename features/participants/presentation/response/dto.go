package response

import (
	"lami/app/config"
	"lami/app/features/participants"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Participant struct {
	ID            int       `json:"participantID" form:"participantID"`
	Name          string    `json:"name" form:"name"`
	Detail        string    `json:"details" form:"details"`
	Image         string    `json:"image" form:"image"`
	HostedBy      string    `json:"hostedby" form:"hostedby"`
	Date          time.Time `json:"date" form:"date"`
	City          string    `json:"city" form:"city"`
	Location      string    `json:"location" form:"location"`
	GrossAmount   int64     `json:"gross_amount"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
	Status        string    `json:"status"`
}

type Payment struct {
	OrderID           string
	TransactionID     string
	PaymentMethod     string
	BillNumber        string
	Bank              string
	GrossAmount       string
	TransactionTime   string
	TransactionExpire string
}

func FromCore(core participants.Core) Participant {
	return Participant{
		ID:            core.ID,
		Name:          core.Event.Name,
		Detail:        core.Event.Detail,
		Image:         core.Event.Image,
		HostedBy:      core.Event.HostedBy,
		Date:          core.Date,
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

func FromMidtransToPayment(resMidtrans *coreapi.ChargeResponse) Payment {
	return Payment{
		OrderID:         resMidtrans.OrderID,
		TransactionID:   resMidtrans.TransactionID,
		PaymentMethod:   config.PaymentBankTransferBCA,
		BillNumber:      resMidtrans.VaNumbers[0].VANumber,
		Bank:            resMidtrans.VaNumbers[0].Bank,
		GrossAmount:     resMidtrans.GrossAmount,
		TransactionTime: resMidtrans.TransactionTime,
	}
}
