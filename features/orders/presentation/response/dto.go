package response

import (
	"lami/app/features/orders"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Order struct {
	Receiver   string        `json:"receiver"`
	Address    string        `json:"address"`
	TotalPrice uint          `json:"totalprice"`
	Status     string        `json:"status"`
	Product    []OrderDetail `json:"product"`
}

type OrderDetail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Qty  uint   `json:"qty"`
}

func FromCore(core orders.Core) Order {
	return Order{
		Receiver:   core.Receiver,
		Address:    core.Address,
		TotalPrice: core.TotalPrice,
		Status:     core.Status,
		Product:    FromCoreDetailList(core.Product),
	}
}

func FromCoreList(core []orders.Core) []Order {
	res := []Order{}
	for v := range core {
		res = append(res, FromCore(core[v]))
	}
	return res
}

func FromCoreDetail(core orders.CoreDetail) OrderDetail {
	return OrderDetail{
		ID:   core.Product.ID,
		Name: core.Product.Name,
		Url:  core.Product.Url,
		Qty:  core.Product.Qty,
	}
}

func FromCoreDetailList(core []orders.CoreDetail) []OrderDetail {
	res := []OrderDetail{}
	for v := range core {
		res = append(res, FromCoreDetail(core[v]))
	}
	return res
}

//	Payments

func FromCoreChargeMidtrans(dataResp coreapi.ChargeResponse) orders.CoreChargeResponse {
	return orders.CoreChargeResponse{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		VAnumbers: orders.VAnumbers{
			BankTransfer: dataResp.VaNumbers[0].Bank,
			VAnumber:     dataResp.VaNumbers[0].VANumber,
		},
		OrderID:  dataResp.OrderID,
		GroosAmt: dataResp.GrossAmount,
	}
}

func FromCoreChargePermata(dataResp coreapi.ChargeResponse) orders.CoreChargeResponsePermata {
	return orders.CoreChargeResponsePermata{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		PermataVaNumber:   dataResp.PermataVaNumber,
		OrderID:           dataResp.OrderID,
		GroosAmt:          dataResp.GrossAmount,
	}
}

func FromCoreChargeMandiri(dataResp coreapi.ChargeResponse) orders.CoreChargeResponseMandiri {
	return orders.CoreChargeResponseMandiri{
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		PaymentType:       dataResp.PaymentType,
		OrderID:           dataResp.OrderID,
		GroosAmt:          dataResp.GrossAmount,
		BillKey:           dataResp.BillKey,
		BillerCode:        dataResp.BillerCode,
	}
}

func FromCoreStatusResponse(dataResp coreapi.TransactionStatusResponse) orders.CoreTransactionStatusResponse {
	return orders.CoreTransactionStatusResponse{
		OrderID:           dataResp.OrderID,
		TransactionTime:   dataResp.TransactionTime,
		TransactionStatus: dataResp.TransactionStatus,
		GrossAmount:       dataResp.GrossAmount,
		PaymentType:       dataResp.PaymentType,
		SettlementTime:    dataResp.SettlementTime,
	}
}
