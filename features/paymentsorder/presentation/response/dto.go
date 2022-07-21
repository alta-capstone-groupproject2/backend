package response

import (
	"lami/app/features/paymentsorder"

	"github.com/midtrans/midtrans-go/coreapi"
)

func FromCoreChargeMidtrans(dataResp coreapi.ChargeResponse) paymentsorder.CoreChargeResponse {
	return paymentsorder.CoreChargeResponse{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		VAnumbers: paymentsorder.VAnumbers{
			BankTransfer: dataResp.VaNumbers[0].Bank,
			VAnumber:     dataResp.VaNumbers[0].VANumber,
		},
		OrderID:  dataResp.OrderID,
		GroosAmt: dataResp.GrossAmount,
	}
}

func FromCoreChargePermata(dataResp coreapi.ChargeResponse) paymentsorder.CoreChargeResponsePermata {
	return paymentsorder.CoreChargeResponsePermata{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		PermataVaNumber:   dataResp.PermataVaNumber,
		OrderID:           dataResp.OrderID,
		GroosAmt:          dataResp.GrossAmount,
	}
}

func FromCoreChargeMandiri(dataResp coreapi.ChargeResponse) paymentsorder.CoreChargeResponseMandiri {
	return paymentsorder.CoreChargeResponseMandiri{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		OrderID:           dataResp.OrderID,
		GroosAmt:          dataResp.GrossAmount,
		BillKey:           dataResp.BillKey,
		BillerCode:        dataResp.BillerCode,
	}
}

func FromCoreStatusResponse(dataResp coreapi.TransactionStatusResponse) paymentsorder.CoreTransactionStatusResponse {
	return paymentsorder.CoreTransactionStatusResponse{
		OrderID:           dataResp.OrderID,
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		GrossAmount:       dataResp.GrossAmount,
		PaymentType:       dataResp.PaymentType,
		SettlementTime:    dataResp.SettlementTime,
	}
}
