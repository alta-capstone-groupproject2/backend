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
	GetAllEvent(limit, offset int, city, name string) (dataEvent []Core, total int64, err error)
	GetEventByID(eventID int) (dataEvent Core, err error)
	InsertEvent(dataReq Core) (err error)
	DeleteEventByID(eventID, userID int) (err error)
	UpdateEventByID(status string, eventID int) (err error)
	GetEventByUserID(userID, limit, offset int) (dataEvent []Core, total int64, err error)
	GetEventSubmission(limit, offset int) (dataEvent []Submission, total int64, err error)
	GetEventSubmissionByID(eventID int) (dataEvent Core, err error)
	GetEventAttendee(eventID, userID int) (urlPDF string, err error)
}

type Data interface {
	SelectData(limit int, offset int, city string, name string) (data []Core, total int64, err error)
	SelectDataByID(eventID int) (dataEvent Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteDataByID(eventID, userID int) (err error)
	UpdateDataByID(status string, eventID, userID int) (err error)
	SelectDataByUserID(userID, limit, offset int) (dataEvent []Core, total int64, err error)
	SelectParticipantData(eventID int) (dataEvent []Participant, err error)
	/** untuk check userID
	**
	** @param id by EventID
	** @return UserID, error
	**
	**/
	CheckValidateUserID(eventID int) (userID int, err error)

	SelectDataSubmission(limit, offset int) (dataEvent []Submission, total int64, err error)
	SelectDataSubmissionByID(eventID int) (dataEvent Core, err error)

	SelectAttendeeData(eventID int) (dataEvent []AttendeesData, err error)
}
