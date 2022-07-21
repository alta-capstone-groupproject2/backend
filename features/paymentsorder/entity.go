package paymentsorder

import _dataOrder "lami/app/features/orders/data"

type Business interface {
	Payments(idUser int) (int, int, error)
}

type Data interface {
	DataPayments(idUser int) (int, int, error)
}

type CoreChargeRequest struct {
	PaymentType        string
	Items              ItemDetails
	Customer           CustomerDetails
	TransactionDetails TransactionDetails
	BankTransfer       BankTransferDetails
	EChannel           EChannelDetail
	Qris               QrisDetails
	GoPay              GopayDetails
	ShoopePay          ShoopePayDetails
	ConvStore          ConvStoreDetails
	Order              _dataOrder.Order
}

type TransactionDetails struct {
	OrderID  string
	GroosAmt int64
}

type ItemDetails struct {
	ID           string
	Name         string
	Price        int64
	Qty          int32
	MerchantName string
}

type CustomerDetails struct {
	FullName string
	Email    string
	Phone    string
}

type EChannelDetail struct {
	BillInfo1 string
	BillInfo2 string
	BillKey   string
}

type BankTransferDetails struct {
	BankName       string
	VANumber       string
	Permata        PermataBankTransferDetail
	LangIDInquiry  string
	LangENInquiry  string
	LangIDPayment  string
	LangENPayment  string
	SubCompanyCode string
	BillInfo1      string
	BillInfo2      string
	BillInfoKey    string
}

type PermataBankTransferDetail struct {
	RecipientName string
}

type QrisDetails struct {
	Acquirer string
}

type GopayDetails struct {
	EnableCallback     bool
	CallbackUrl        string
	AccountID          string
	PaymentOptionToken string
}

type ShoopePayDetails struct {
	CallbackUrl string
}

type ConvStoreDetails struct {
	Store string
}

type VAnumbers struct {
	BankTransfer string
	VAnumber     string
}

//ChargeResponse : CoreAPI charge response struct when calling Midtrans API
type CoreChargeResponse struct {
	TransactionTime   string
	TransactionStatus string
	PaymentType       string
	VAnumbers         VAnumbers
	OrderID           string
	GroosAmt          string
}

//ChargeResponsePermata : CoreAPI charge response struct when calling Midtrans API
type CoreChargeResponsePermata struct {
	TransactionTime   string
	TransactionStatus string
	PaymentType       string
	PermataVaNumber   string
	OrderID           string
	GroosAmt          string
}

//ChargeResponseMandiri : CoreAPI charge response struct when calling Midtrans API
type CoreChargeResponseMandiri struct {
	TransactionTime   string
	TransactionStatus string
	PaymentType       string
	OrderID           string
	GroosAmt          string
	BillKey           string
	BillerCode        string
}

//TransactionStatusResponse : Status transaction response struct
type CoreTransactionStatusResponse struct {
	OrderID           string
	TransactionTime   string
	TransactionStatus string
	Bank              string
	GrossAmount       string
	PaymentType       string
	SettlementTime    string
}
