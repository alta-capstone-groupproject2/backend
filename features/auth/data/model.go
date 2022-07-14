package data

import (
	"lami/app/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

//DTO

func (data *User) toCore() auth.Core {
	return auth.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func authCore(core auth.Core) User {
	return User{
		Email:    core.Email,
		Password: core.Password,
	}
}
