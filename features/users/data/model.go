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

func (data *User) toCore() users.Core {
	return users.Core{
		ID:        int(data.ID),
		Image:     data.Image,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Role: users.Role{
			RoleName: data.Role.RoleName,
		},
	}

}

// func toCoreList(data []User) []users.Core {
// 	result := []users.Core{}
// 	for key := range data {
// 		result = append(result, data[key].toCore())
// 	}
// 	return result
// }

func fromCore(core users.Core) User {
	return User{
		Image:    core.Image,
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
		RoleID:   core.RoleID,
	}
}

func toCore(data User) users.Core {
	return data.toCore()
}
