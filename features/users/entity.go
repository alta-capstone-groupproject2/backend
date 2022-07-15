package users

import (
	"mime/multipart"
	"time"
)

type Core struct {
	ID          int
	Name        string
	Email       string
	Password    string
	Image       string
	StoreName   string
	Phone       string
	Owner       string
	City        string
	Address     string
	Document    string
	RoleID      int
	StoreStatus string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Role        Role
}

type Role struct {
	ID       int
	RoleName string
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, total int64, err error)
	GetDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteData(param int) (err error)
	UpdateData(dataReq Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) (err error)
	UpgradeAccount(dataReq Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) error
	UpdateStatusUser(status string, id int) error
}

type Data interface {
	SelectData(limit, offset int) (data []Core, total int64, err error)
	SelectDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (err error)
	DeleteData(param int) (err error)
	UpdateData(dataReq map[string]interface{}, id int) (err error)
	InsertStoreData(dataReq Core, id int) error
	UpdateAccountRole(status string, id int) error
}
