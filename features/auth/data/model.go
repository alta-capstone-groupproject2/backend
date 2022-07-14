package data

import (
	"lami/app/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Image       string `json:"image"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RoleID      int    `json:"role_id"`
	storeName   string `json:"store_name"`
	phone       string `json:"phone"`
	storeOwner  string `json:"store_owner"`
	city        string `json:"city"`
	Address     string `json:"address"`
	Document    string `json:"document"`
	StoreStatus string `json:"store_status"`
	Role        Role
}

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"role_name"`
}

//DTO

func (data *User) toCore() auth.Core {
	return auth.Core{
		ID:        int(data.ID),
		Name:      data.Name,
		Email:     data.Email,
		Role:      data.Role.RoleName,
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
