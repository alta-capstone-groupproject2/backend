package request

import (
	"lami/app/features/orders"
	"strconv"
	"time"

	"lami/app/features/orders/data"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type Order struct {
	CartID      []int   `json:"cart_id" form:"cart_id"`
	UserID      int     `json:"user_id" form:"user_id"`
	Receiver    string  `json:"receiver" form:"receiver"`
	PhoneNumber string  `json:"phone" form:"phone"`
	Address     string  `json:"address" form:"address"`
	TotalPrice  float32 `json:"totalprice" form:"totalprice"`
	Status      string  `json:"status" form:"status"`
}

func ToCore(orderReq Order) orders.Core {
	var dataCartID []int
	for key := range orderReq.CartID {
		dataCartID = append(dataCartID, key)
	}

	return orders.Core{
		UserID:      orderReq.UserID,
		CartID:      dataCartID,
		Receiver:    orderReq.Receiver,
		PhoneNumber: orderReq.PhoneNumber,
		Address:     orderReq.Address,
	}

}

//	Payment

type ChargeRequest struct {
	PaymentType        coreapi.CoreapiPaymentType
	TransactionDetails midtrans.TransactionDetails
	BankTransfer       *coreapi.BankTransferDetails
	EChannelDetails    coreapi.EChannelDetail
	Order              data.Order
}

func ToCoreMidtransBank(dataReq ChargeRequest) coreapi.ChargeReq {

	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dataReq.TransactionDetails.OrderID,
			GrossAmt: dataReq.TransactionDetails.GrossAmt,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: dataReq.BankTransfer.Bank,
		},
	}
}

func ToCoreMidtransPermata(dataReq ChargeRequest) coreapi.ChargeReq {

	return coreapi.ChargeReq{
		PaymentType: "bank_transfer",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dataReq.TransactionDetails.OrderID,
			GrossAmt: dataReq.TransactionDetails.GrossAmt,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: dataReq.BankTransfer.Bank,
			Permata: &coreapi.PermataBankTransferDetail{
				RecipientName: dataReq.BankTransfer.Permata.RecipientName,
			},
		},
	}
}

func ToCoreMidtransMandiri(dataReq ChargeRequest) coreapi.ChargeReq {
	return coreapi.ChargeReq{
		PaymentType: "echannel",
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  dataReq.TransactionDetails.OrderID,
			GrossAmt: dataReq.TransactionDetails.GrossAmt,
		},
		EChannel: &coreapi.EChannelDetail{
			BillInfo1: dataReq.EChannelDetails.BillInfo1,
			BillInfo2: dataReq.EChannelDetails.BillInfo2,
		},
	}
}

func ToCoreMidtransEMoney(dataReq ChargeRequest) coreapi.ChargeReq {
	return coreapi.ChargeReq{}
}

func Random() string {
	time.Sleep(1000000 * time.Microsecond)
	return strconv.FormatInt(time.Now().Unix(), 10)
}
