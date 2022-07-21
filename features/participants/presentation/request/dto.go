package request

import (
	"lami/app/features/participants"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Participant struct {
	UserID        int `json:"userID" form:"userID"`
	EventID       int `json:"eventID" form:"eventID"`
	OrderID       string
	GrossAmount   int64
	PaymentMethod string `json:"paymentMethod" form:"paymentMethod"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
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
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: "bca",
		},
	}
}
