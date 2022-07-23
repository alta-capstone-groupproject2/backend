package participants

import (
	"lami/app/features/events/data"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Core struct {
	ID            int
	UserID        int
	EventID       int
	OrderID       string
	Date          time.Time
	GrossAmount   int64
	PaymentMethod string
	TransactionID string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Event         Event
}

type Event struct {
	ID        int
	Image     string
	Document  string
	Name      string
	HostedBy  string
	Phone     string
	Date      time.Time
	City      string
	Location  string
	Detail    string
	Price     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment struct {
	OrderID           string
	TransactionID     string
	PaymentMethod     string
	BillNumber        string
	Bank              string
	GrossAmount       int64
	TransactionTime   time.Time
	TransactionExpire time.Time
}

type Business interface {
	AddParticipant(data Core) error
	GetAllEventbyParticipant(userID int) (data []Core, err error)
	DeleteParticipant(param, userID int) error

	//Payment Event
	GrossAmountEvent(id int) (GrossAmount int64, err error)
	GetDetailPayment(limit, page, userID int) ([]Core, int64, error)
	CreatePaymentBankTransfer(coreapi.ChargeReq, Core) (*coreapi.ChargeResponse, error)
	CheckStatusPayment(orderID string) (*coreapi.TransactionStatusResponse, error)
	PaymentWebHook(orderID, status string) error
}

type Data interface {
	SelectDataByID(id int) (response data.Event, err error)
	AddData(data Core) error
	SelectDataEvent(eventID int) (data []Core, err error)
	DeleteData(param, userID int) error

	//validasi join event
	SelectValidasi(userID, eventID int) bool

	//Payment Event Data
	UpdateDataPayment(*coreapi.ChargeResponse, Core) error
	SelectPayment(orderID string) (Core, error)
	SelectPaymentList(limit, page, userID int) ([]Core, int64, error)
	CreateDataPayment(coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	CheckDataStatusPayment(orderID string) (*coreapi.TransactionStatusResponse, error)
	PaymentDataWebHook(data Core) error
}
