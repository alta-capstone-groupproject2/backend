package response

import (
	"lami/app/features/users"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Image     string    `json:"image"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserStore struct {
	ID         int       `json:"id" form:"id"`
	Image      string    `json:"image" form:"image"`
	Name       string    `json:"name" form:"name"`
	Email      string    `json:"email" form:"email"`
	StoreName  string    `json:"store_name"`
	Phone      string    `json:"phone"`
	Role       string    `json:"role"`
	StoreOwner string    `json:"store_owner"`
	City       string    `json:"city"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
}

func FromCore(data users.Core) User {
	return User{
		ID:        data.ID,
		Image:     data.Image,
		Name:      data.Name,
		Email:     data.Email,
		Role:      data.Role.RoleName,
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

func UserStoreFromCore(data users.Core) UserStore {
	return UserStore{
		ID:         data.ID,
		Image:      data.Image,
		Name:       data.Name,
		Email:      data.Email,
		CreatedAt:  data.CreatedAt,
		StoreName:  data.StoreName,
		Phone:      data.Phone,
		StoreOwner: data.StoreOwner,
		City:       data.City,
	}
}

func UserStoreFromCoreList(data []users.Core) []UserStore {
	result := []UserStore{}
	for key := range data {
		result = append(result, UserStoreFromCore(data[key]))
	}
	return result
}
