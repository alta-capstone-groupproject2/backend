package users

import (
	"mime/multipart"
	"time"
)

type Core struct {
	ID         int
	Name       string
	Email      string
	Password   string
	Image      string
	StoreName  string
	Phone      string
	StoreOwner string
	City       string
	RoleID     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Role       Role
}

type Role struct {
	ID       int
	RoleName string
}

type Business interface {
	// GetAllData(limit, offset int) (data []Core, err error)
	GetDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteData(param int) (err error)
	UpdateData(dataReq Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) (err error)
}

type Data interface {
	// SelectData(limit, offset int) (data []Core, err error)
	SelectDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteData(param int) (err error)
	UpdateData(dataReq map[string]interface{}, id int) (err error)
}
