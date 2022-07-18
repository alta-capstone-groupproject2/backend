package data

import (
	"lami/app/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Image       string `json:"image"`
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	RoleID      int    `json:"role_id"`
	StoreName   string `json:"store_name"`
	Phone       string `json:"phone"`
	StoreOwner  string `json:"store_owner"`
	City        string `json:"city"`
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

func (data *User) toCore() users.Core {
	return users.Core{
		ID:          int(data.ID),
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		Image:       data.Image,
		StoreName:   data.StoreName,
		Phone:       data.Phone,
		Owner:       data.StoreOwner,
		City:        data.City,
		Address:     data.Address,
		Document:    data.Document,
		RoleID:      data.RoleID,
		StoreStatus: data.StoreStatus,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Role:        users.Role{RoleName: data.Role.RoleName},
	}

}

func fromCore(core users.Core) User {
	return User{
		Image:       core.Image,
		Name:        core.Name,
		Email:       core.Email,
		Password:    core.Password,
		RoleID:      core.RoleID,
		StoreName:   core.StoreName,
		StoreStatus: core.StoreStatus,
		StoreOwner:  core.Owner,
		Phone:       core.Phone,
		City:        core.City,
		Address:     core.Address,
		Document:    core.Document,
	}
}

func StiretoCore(data User) users.Core {
	return data.toCore()
}

func toCore(data User) users.Core {
	return data.toCore()
}

func ToCoreList(data []User) []users.Core {
	result := []users.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}
