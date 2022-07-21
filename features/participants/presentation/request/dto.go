package request

import (
	"lami/app/features/participants"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Participant struct {
	ID            int
	UserID        int `json:"userID" form:"userID"`
	EventID       int `json:"eventID" form:"eventID"`
	OrderID       string
	GrossAmount   int64
	PaymentMethod string `json:"paymentMethod" form:"paymentMethod"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		ID:            partReq.ID,
		UserID:        partReq.UserID,
		EventID:       partReq.EventID,
		OrderID:       partReq.OrderID,
		GrossAmount:   partReq.GrossAmount,
		PaymentMethod: partReq.PaymentMethod,
	}
}

func ToCoreMidtrans(req participants.Core) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.GrossAmount,
		},
	}
}

type MidtransHookRequest struct {
	TransactionTime   string `form:"transaction_time" json:"transaction_time"`
	TransactionStatus string `form:"transaction_status" json:"transaction_status"`
	OrderID           string `form:"order_id" json:"order_id"`
	MerchantID        string `form:"merchant_id" json:"merchant_id"`
	GrossAmount       string `form:"gross_amount" json:"gross_amount"`
	FraudStatus       string `form:"fraud_status" json:"fraud_status"`
	Currency          string `form:"currency" json:"currency"`
}
