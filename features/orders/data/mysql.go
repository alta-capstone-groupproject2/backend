package data

import (
	"fmt"
	"lami/app/features/orders"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	db *gorm.DB
}

// AddDataOrder implements orders.Data
func (repo *mysqlOrderRepository) AddDataOrder(dataReq orders.Core, idUser int) (int, error) {
	//	Check for qty
	dataCart := CartCheck{}
	fmt.Println("dataCart:", dataCart)
	checkStock := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	fmt.Println("checkStock:", checkStock)
	fmt.Println("dataCart after checkStock:", dataCart)
	if checkStock.Error != nil {
		return -2, checkStock.Error
	}



}

func NewOrderRepository(conn *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		db: conn,
	}
}
