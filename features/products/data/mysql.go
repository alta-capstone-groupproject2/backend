package data

import (
	"fmt" 
	"errors"
	"lami/app/features/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// SelectDataRating implements product.Data
func (repo *mysqlProductRepository) SelectDataRating(idProduct int) (product.CoreRating, error) {
	dataRating := Rating{}

	res := repo.db.Preload("Product").Where("product_id = ?", idProduct).Find(&dataRating)
	fmt.Println("res rating from mysql.go:", res)
	if res.Error != nil {
		return product.CoreRating{}, res.Error
	}

	return ToCoreRating(dataRating), nil
}

// SelectDataMyProduct implements product.Data
func (repo *mysqlProductRepository) SelectDataMyProduct(idUser int) ([]product.Core, error) {
	dataMyProduct := []Product{}

	res := repo.db.Preload("User").Where("user_id = ?", idUser).Find(&dataMyProduct)
	fmt.Println("res from mysql.go", res)
	if res.Error != nil {
		return []product.Core{}, res.Error
	}

	return ToCoreMyProduct(dataMyProduct), nil
}

// SelectDataProductbyIDProduct implements product.Data
func (repo *mysqlProductRepository) SelectDataProductbyIDProduct(idProduct int) (product.Core, error) {
	dataProduct := Product{}

	// var idUser int
	// idUserJoin := repo.db.Raw("SELECT user_id FROM products WHERE id = ?", idProduct).Scan(&idUser)
	// fmt.Println(idUserJoin)
	// fmt.Println("idUser:", idUser)

	res := repo.db.Preload("User").Where("id = ?", idProduct).Find(&dataProduct)
	// res := repo.db.Find(&dataProduct)
	fmt.Println("res from mysql.go:", res)
	if res.Error != nil {
		return product.Core{}, res.Error
	}

	return ToCore(dataProduct), res.Error
}

// DeleteDataProduct implements product.Data
func (repo *mysqlProductRepository) DeleteDataProduct(idProduct int, idUser int) error {
	dataProduct := Product{}
	res := repo.db.Where("user_id = ?", idUser).Delete(&dataProduct, idProduct)
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// UpdateDataProduct implements product.Data
func (repo *mysqlProductRepository) UpdateDataProduct(dataReq product.Core, idProduct, idUser int) error {

	model := Product{}
	model.ID = uint(idProduct)
	err := repo.db.Model(model).Where("user_id = ?", idUser).Updates(dataReq)
	if err.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if err != nil {
		return err.Error
	}

	return nil
}

// AddDataProduct implements product.Data
func (repo *mysqlProductRepository) AddDataProduct(dataReq product.Core) (int, error) {

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

// AddProductDataRating implements product.Data
func (repo *mysqlProductRepository) AddProductDataRating(dataReq product.CoreRating) (int, error) {

	model := fromCoreRating(dataReq)

	res := repo.db.Create(&model)
	if res.Error != nil {
		return -1, res.Error
	}
	if res.RowsAffected == 0 {
		return 0, errors.New("failed")
	}

	return int(res.RowsAffected), nil
}

func NewProductRepository(conn *gorm.DB) product.Data {
	return &mysqlProductRepository{
		db: conn,
	}
}
