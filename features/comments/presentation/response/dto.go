package response

import (
	"lami/app/features/comments"
	"time"
)

type Comment struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	Avatar    string    `json:"avatar" form:"avatar"`
	Comment   string    `json:"comment" form:"comment"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func FromCore(data comments.Core) Comment {
	return Comment{
		ID:        data.ID,
		Name:      data.User.Name,
		Avatar:    data.User.Image,
		Comment:   data.Comment,
		UpdatedAt: data.UpdatedAt,
	}
}

func FromCoreList(data []comments.Core) []Comment {
	result := []Comment{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
