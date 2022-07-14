package request

import "lami/app/features/users"

type User struct {
	URL      string `json:"url" form:"url"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq User) users.Core {
	userCore := users.Core{
		URL:      userReq.URL,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
