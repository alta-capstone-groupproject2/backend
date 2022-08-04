package comments

import (
	"time"
)

type Core struct {
	ID        int
	EventID   int
	UserID    int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User
	Event     Event
}

type Event struct {
	ID        int
	Image     string
	Document  string
	Name      string
	HostedBy  string
	Phone     string
	StartDate time.Time
	EndDate   time.Time
	City      string
	Location  string
	Detail    string
	Price     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	Image       string
	StoreName   string
	Phone       string
	Owner       string
	City        string
	Address     string
	Document    string
	RoleID      int
	StoreStatus string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Business interface {
	AddComment(dataEvent Core) (err error)
	GetCommentByIdEvent(limit, offset, event_id int) (data []Core, count int64, err error)
}

type Data interface {
	Insert(dataEvent Core) (err error)
	GetComment(limit, offset, event_id int) (data []Core, count int64, err error)
}
