package participants

import (
	"lami/app/features/events/data"
	"time"
)

type Core struct {
	ID            int
	UserID        int
	EventID       int
	OrderID       string
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
	// CreatePaymentBankTransferBCA(payment Core) (Payment, error)
}

type Data interface {
	SelectDataByID(id int) (response data.Event, err error)
	AddData(data Core) error
	SelectDataEvent(userID int) (data []Core, err error)
	DeleteData(param, userID int) error
}
