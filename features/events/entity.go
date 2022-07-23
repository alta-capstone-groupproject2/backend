package events

import (
	"time"
)

type Core struct {
	ID            int
	Image         string
	Document      string
	Name          string
	HostedBy      string
	Phone         string
	StartDate     time.Time
	EndDate       time.Time
	City          string
	Location      string
	Detail        string
	Price         int
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int
	Participant   []Participant
	AttendeesData []AttendeesData
}

type Participant struct {
	ID    int
	Name  string
	Image string
}

type Submission struct {
	ID        int
	Name      string
	UserName  string
	City      string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

type AttendeesData struct {
	Num     int
	Name    string
	Email   string
	City    string
	Present string
}

type Business interface {
	GetAllEvent(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	GetEventByID(param int) (data Core, err error)
	InsertEvent(dataReq Core) (err error)
	DeleteEventByID(id int, userId int) (err error)
	UpdateEventByID(status string, id int) (err error)
	GetEventByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
	GetEventSubmission(limit, offset int) (data []Submission, total int64, err error)
	GetEventSubmissionByID(id int) (data Core, err error)
	GetEventAttendee(id, userID int) (urlPDF string, err error)
}

type Data interface {
	SelectData(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	SelectDataByID(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteDataByID(id int, userId int) (err error)
	UpdateDataByID(status string, id int, userId int) (err error)
	SelectDataByUserID(id_user, limit, offset int) (data []Core, total int64, err error)
	SelectParticipantData(id_event int) (data []Participant, err error)
	/** untuk check userID
	**
	** @param id by EventID
	** @return UserID, error
	**
	**/
	CheckUserID(id int) (userID int, err error)

	SelectDataSubmission(limit, offset int) (data []Submission, total int64, err error)
	SelectDataSubmissionByID(id int) (data Core, err error)

	SelectAttendeeData(id_event int) (data []AttendeesData, err error)
}
