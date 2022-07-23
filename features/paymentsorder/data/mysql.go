package data

import (
	"lami/app/features/paymentsorder"

	"gorm.io/gorm"
)

type mysqlPaymentsRepository struct {
	db *gorm.DB
}

// InsertDataPayments implements paymentsorder.Data
func (repo *mysqlPaymentsRepository) DataPayments(idUser int) (int, int, error) {
	var idOrder, totalprice int
	residOrder := repo.db.Raw("SELECT id FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", idUser).Scan(&idOrder)
	if residOrder.Error != nil {
		return -1, -1, residOrder.Error
	}

	resTotalPrice := repo.db.Raw("SELECT total_price FROM orders WHERE user_id = ? ORDER BY id DESC LIMIT 1", idUser).Scan(&totalprice)
	if resTotalPrice.Error != nil {
		return -1, -1, resTotalPrice.Error
	}

	return idOrder, totalprice, nil
}

func NewPaymentsRepository(conn *gorm.DB) paymentsorder.Data {
	return &mysqlPaymentsRepository{
		db: conn,
	}
}