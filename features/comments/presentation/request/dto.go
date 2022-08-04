package request

import (
	"lami/app/features/comments"
)

type Comment struct {
	EventID int    `json:"eventID" form:"eventID"`
	Comment string `json:"comment" form:"comment"`
}

func ToCore(req Comment) comments.Core {
	commentCore := comments.Core{
		EventID: req.EventID,
		Comment: req.Comment,
	}
	return commentCore
}
