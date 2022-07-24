package orders

// import "lami/app/features/products/data"

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

type CoreResponse struct {
	Receiver   string
	Address    string
	TotalPrice uint
	Status     string
	Product    []Product
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

	PaymentsOrderID(idUser int) (int, error)
	PaymentGrossAmount(idUser int) (int, error)
}

type Data interface {
	SelectDataHistoryOrder(idUser int) ([]Core, error)

	UpdateStockOnProductPlusCountTotalPrice(dataReq Core, idUser int) (int, error)
	DeleteDataCart(dataReq Core, idUser int) error
	AddDataOrder(dataReq Core, idUser int, total int) (int64, error)
	AddDataOrderDetail(dataReq Core, row int64, idUser int) error

	DataPaymentsOrderID(idUser int) (int, error)
	DataPaymentsGrossAmount(idUser int) (int, error)
}

//	Payments
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
	Order              Core
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
