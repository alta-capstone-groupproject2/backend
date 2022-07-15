package events

import (
	"time"
)

type Core struct {
	ID          int
	Image       string
	Document    string
	Name        string
	HostedBy    string
	Phone       string
	Date        time.Time
	City        string
	Location    string
	Detail      string
	Price       int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IDUser      int
	Participant []Participant
}

type Participant struct {
	ID    int
	Name  string
	Image string
}

type Business interface {
	GetAllEvent(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	GetEventByID(param int) (data Core, err error)
	InsertEvent(dataReq Core) (err error)
	DeleteEventByID(id int, userId int) (err error)
	UpdateEventByID(dataReq Core, id int, userId int) (err error)
	GetEventByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
}

type Data interface {
	SelectData(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	SelectDataByID(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteDataByID(id int, userId int) (err error)
	UpdateDataByID(dataReq map[string]interface{}, id int, userId int) (err error)
	SelectDataByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
	SelectParticipantData(id_event int) (data []Participant, err error)
}
