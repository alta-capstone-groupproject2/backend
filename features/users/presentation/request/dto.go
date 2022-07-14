package request

import "lami/app/features/users"

type User struct {
	Image    string `json:"image" form:"image"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq User) users.Core {
	userCore := users.Core{
		Image:    userReq.Image,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
