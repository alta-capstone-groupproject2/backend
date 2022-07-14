package request

import "lami/app/features/participants"

type Participant struct {
	UserID  int `json:"id_user" form:"id_user"`
	EventID int `json:"id_event" form:"id_event"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		UserID:  partReq.UserID,
		EventID: partReq.EventID,
	}
}
