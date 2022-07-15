package data

import (
	"errors"
	"lami/app/features/carts"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	db *gorm.DB
}

// DeleteDataCart implements cart.Data
func (repo *mysqlCartRepository) DeleteDataCart(idCart, idUser int) error {
	dataCart := Cart{}
	res := repo.db.Where("id = ? AND user_id = ?", idCart, idUser).Delete(&dataCart)
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// UpdateDataCart implements cart.Data
func (repo *mysqlCartRepository) UpdateDataCart(dataReq cart.Core, idCart int) error {
	model := Cart{}
	err := repo.db.Model(model).Where("id = ?", idCart).Updates(dataReq)
	if err.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if err != nil {
		return err.Error
	}

	return nil
}

// SelectDataCart implements cart.Data
func (repo *mysqlCartRepository) SelectDataCart(idUser int) ([]cart.Core, error) {
	dataCart := []Cart{}

	result := repo.db.Preload("Product").Where("user_id = ?", idUser).Find(&dataCart)
	if result.Error != nil {
		return []cart.Core{}, result.Error
	}

	return ToCoreList(dataCart), nil
}

// AddDataCart implements cart.Data
func (repo *mysqlCartRepository) AddDataCart(dataReq cart.Core) (int, error) {
	model := fromCore(dataReq)

	res := repo.db.Create(&model)
	if res.Error != nil {
		return -1, res.Error
	}
	if res.RowsAffected == 0 {
		return 0, errors.New("failed")
	}

	return int(res.RowsAffected), nil
}

func NewCartRepository(conn *gorm.DB) cart.Data {
	return &mysqlCartRepository{
		db: conn,
	}
}
