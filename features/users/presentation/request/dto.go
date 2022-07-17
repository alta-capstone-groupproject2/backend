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

type Store struct {
	StoreName   string `json:"storeName" form:"storeName"`
	Phone       string `json:"phone" form:"phone"`
	Owner       string `json:"owner" form:"owner"`
	City        string `json:"city" form:"city"`
	Address     string `json:"address" form:"address"`
	Document    string
	StoreStatus string `json:"status" form:"status"`
}

type Gmail struct {
	Gmail string `json:"email" form:"email"`
}

func StoreToCore(userReq Store) users.Core {
	userCore := users.Core{
		StoreName: userReq.StoreName,
		Phone:     userReq.Phone,
		Owner:     userReq.Owner,
		City:      userReq.City,
		Address:   userReq.Address,
	}
	return userCore
}
