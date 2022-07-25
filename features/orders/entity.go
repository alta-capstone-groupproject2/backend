package orders

import (
	"lami/app/features/users/data"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Core struct {
	ID          int
	CartID      []int
	UserID      int
	Receiver    string
	PhoneNumber string
	TotalPrice  uint
	Address     string
	Status      string
	Product     []CoreDetail
}

type CoreDetail struct {
	ID        int
	OrderID   int
	ProductID int
	Qty       uint
	Product   Product
}

type Product struct {
	ID   int
	Name string
	Url  string
	Qty  uint
}

type Business interface {
	Order(dataReq Core, idUser int) (int, error)
	SelectHistoryOrder(idUser int) ([]Core, error)

	TypeBank(grossamount int64, typename string, idOrder int) (coreapi.ChargeReq, string, error)
	RequestChargeBank(dataCore coreapi.ChargeReq, typename string) (coreapi.ChargeReq, error)
	PaymentsOrderID(idUser int) (int, error)
	PaymentGrossAmount(idUser int) (int, error)
	UpdateStatus(idOrder int) error
}

type Data interface {
	SelectDataHistoryOrder(idUser int) ([]Core, error)

	UpdateStockOnProductPlusCountTotalPrice(dataReq Core, idUser int) (int, error)
	DeleteDataCart(dataReq Core, idUser int) error
	AddDataOrder(dataReq Core, idUser int, total int) (int64, error)
	AddDataOrderDetail(dataReq Core, row int64, idUser int) error

	DataPaymentsOrderID(idUser int) (int, error)
	DataPaymentsGrossAmount(idUser int) (int, error)
	UpdateDataStatus(idOrder int) error

	SelectUser(id int) (response data.User, err error)
}

//	Payments
type CoreChargeRequest struct {
	PaymentType        string
	TransactionDetails TransactionDetails
	BankTransfer       BankTransferDetails
	EChannel           EChannelDetail
	Order              Core
}

type TransactionDetails struct {
	OrderID  string
	GroosAmt int64
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
