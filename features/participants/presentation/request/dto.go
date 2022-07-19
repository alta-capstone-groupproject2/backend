package request

import "lami/app/features/participants"

type Participant struct {
	UserID        int    `json:"userID" form:"userID"`
	EventID       int    `json:"eventID" form:"eventID"`
	PaymentMethod string `json:"paymentMethod" form:"paymentMethod"`
}

func ToCore(partReq Participant) participants.Core {
	return participants.Core{
		UserID:        partReq.UserID,
		EventID:       partReq.EventID,
		PaymentMethod: partReq.PaymentMethod,
	}
}
