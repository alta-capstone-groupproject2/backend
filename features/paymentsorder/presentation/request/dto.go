package request

import (
	"strconv"
	"time"

	_dataOrder "lami/app/features/orders/data"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type ChargeRequest struct {
	PaymentType        coreapi.CoreapiPaymentType
	TransactionDetails midtrans.TransactionDetails
	BankTransfer       *coreapi.BankTransferDetails
	EChannelDetails    coreapi.EChannelDetail
	Order              _dataOrder.Order
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
