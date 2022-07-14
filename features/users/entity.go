package users

import (
	"time"
)

type Core struct{
	ID 			int
	Name 		string
	Email 		string
	Password 	string
	URL			string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Business interface {
	// GetAllData(limit, offset int) (data []Core, err error)
	GetDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (row int, err error)
	DeleteData(param int) (row int, err error)
	UpdateData(dataReq Core, id int) (row int, err error)
}

type Data interface {
	// SelectData(limit, offset int) (data []Core, err error)
	SelectDataById(param int) (data Core, err error)
	InsertData(dataReq Core) (row int, err error)
	DeleteData(param int) (row int, err error)
	UpdateData(dataReq map[string]interface{}, id int) (row int, err error)
}
