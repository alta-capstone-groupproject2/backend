package request

import "lami/app/features/participants"

type Participant struct {
	UserID  int `json:"userID" form:"userID"`
	EventID int `json:"eventID" form:"eventID"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		UserID:  partReq.UserID,
		EventID: partReq.EventID,
	}
}
