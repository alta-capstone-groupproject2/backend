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

type PaymentDetails struct {
	ID            int       `json:"participantID" form:"participantID"`
	Name          string    `json:"name" form:"name"`
	Date          time.Time `json:"date" form:"date"`
	City          string    `json:"city" form:"city"`
	OrderID       string    `json:"orderID" form:"orderID"`
	GrossAmount   int64     `json:"grossAmount"`
	PaymentMethod string    `json:"paymentMethod"`
	TransactionID string    `json:"transactionID"`
	Status        string    `json:"status"`
}

type Payment struct {
	OrderID           string    `json:"orderID" form:"orderID"`
	TransactionID     string    `json:"transactionID" form:"transactionID"`
	PaymentMethod     string    `json:"paymentMethod" form:"paymentMethod"`
	BillNumber        string    `json:"billNumber" form:"billNumber"`
	Bank              string    `json:"bank" form:"bank"`
	GrossAmount       string    `json:"grossAmount" form:"grossAmount"`
	TransactionTime   time.Time `json:"transactionTime" form:"transactionTime"`
	TransactionExpire time.Time `json:"transactionExpired" form:"transactionExpired"`
	TransactionStatus string    `json:"transactionStatus" form:"transactionStatus"`
}

func FromCore(core participants.Core) Participant {
	return Participant{
		ID:            core.ID,
		Name:          core.Event.Name,
		Detail:        core.Event.Detail,
		Image:         core.Event.Image,
		HostedBy:      core.Event.HostedBy,
		Date:          core.CreatedAt,
		City:          core.Event.City,
		Location:      core.Event.Location,
		GrossAmount:   core.GrossAmount,
		PaymentMethod: core.PaymentMethod,
		TransactionID: core.TransactionID,
		Status:        core.Status,
	}
}

func FromCoreToDetailPayment(core participants.Core) PaymentDetails {
	return PaymentDetails{
		ID:            core.ID,
		Name:          core.Event.Name,
		Date:          core.Date,
		City:          core.Event.City,
		OrderID:       core.OrderID,
		GrossAmount:   core.GrossAmount,
		PaymentMethod: core.PaymentMethod,
		TransactionID: core.TransactionID,
		Status:        core.Status,
	}
}

func FromCoreToDetailPaymentList(data []participants.Core) []PaymentDetails {
	result := []PaymentDetails{}
	for key := range data {
		result = append(result, FromCoreToDetailPayment(data[key]))
	}
	return result
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
		OrderID:           resMidtrans.OrderID,
		TransactionID:     resMidtrans.TransactionID,
		PaymentMethod:     config.PaymentBankTransferBCA,
		BillNumber:        resMidtrans.VaNumbers[0].VANumber,
		Bank:              resMidtrans.VaNumbers[0].Bank,
		GrossAmount:       resMidtrans.GrossAmount,
		TransactionStatus: resMidtrans.TransactionStatus,
	}
}

func FromMidtransToStatusPayment(resMidtrans *coreapi.TransactionStatusResponse) Payment {
	return Payment{
		OrderID:           resMidtrans.OrderID,
		TransactionID:     resMidtrans.TransactionID,
		PaymentMethod:     config.PaymentBankTransferBCA,
		BillNumber:        resMidtrans.VaNumbers[0].VANumber,
		Bank:              resMidtrans.VaNumbers[0].Bank,
		GrossAmount:       resMidtrans.GrossAmount,
		TransactionStatus: resMidtrans.TransactionStatus,
	}
}
