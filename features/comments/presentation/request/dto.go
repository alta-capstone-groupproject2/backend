package request

import (
	"lami/app/features/comments"
)

type Comment struct {
	EventID int    `json:"id_event" form:"id_event"`
	UserID  int    `json:"id_user" form:"id_user"`
	Comment string `json:"comment" form:"comment"`
}

func ToCore(req Comment) comments.Core {
	commentCore := comments.Core{
		EventID: req.EventID,
		UserID:  req.UserID,
		Comment: req.Comment,
	}
	return commentCore
}
