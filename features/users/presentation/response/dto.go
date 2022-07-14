package response

import (
	"lami/app/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id" form:"id"`
	URL       string    `json:"url" form:"url"`
	Name      string    `json:"name" form:"name"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		URL:       data.URL,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}
}

func FromCoreList(data []users.Core) []User {
	result := []User{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
