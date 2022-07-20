package data

import (
	"errors"
	"fmt"
	"lami/app/features/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	db *gorm.DB
}

// SelectProductList implements product.Data
func (repo *mysqlProductRepository) SelectProductList(limit int, page int, city string, name string) ([]product.Core, int64, error) {
	var dataProduct []Product
	var count int64
	res := repo.db.Order("id asc").Where("city LIKE ? and name LIKE ?", "%"+city+"%", "%"+name+"%").Limit(limit).Offset(page).Find(&dataProduct).Count(&count)
	if res.Error != nil {
		return []product.Core{}, 0, res.Error
	}

	return ToCoreListProductList(dataProduct), count, res.Error
}

// SelectDataRating implements product.Data
func (repo *mysqlProductRepository) SelectDataRating(idProduct int) ([]product.CoreRating, error) {
	dataRating := []Rating{}

	res := repo.db.Preload("Product").Where("product_id = ?", idProduct).Find(&dataRating)
	fmt.Println("res rating from mysql.go:", res)
	if res.Error != nil {
		return []product.CoreRating{}, res.Error
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

	res := repo.db.Preload("User").Where("id = ?", idProduct).Find(&dataProduct)
	fmt.Println("res from mysql.go:", res)
	if res.Error != nil {
		return product.Core{}, res.Error
	}

	//	Get Rating from rating database
	rows, err := repo.db.Model(&Rating{}).Where("product_id = ?", idProduct).Select("rating").Rows()
	defer rows.Close()
	if err != nil {
		return product.Core{}, res.Error
	}
	var meanrating []float64
	for rows.Next() {
		var data float64
		if errRows := rows.Scan(&data); errRows != nil {
			panic(errRows)
		}
		meanrating = append(meanrating, data)
	}

	var sum float64
	for i := 0; i < len(meanrating); i++ {
		sum = sum + float64(meanrating[i])
	}
	dataProduct.MeanRating = (sum) / float64(len(meanrating))

	//	Update rating in product database
	updateMeanRating := repo.db.Raw("UPDATE products SET mean_rating = ? WHERE id = ?", dataProduct.MeanRating, idProduct)
	if updateMeanRating.Error != nil {
		return product.Core{}, updateMeanRating.Error
	}

	return ToCorebyProductID(dataProduct), res.Error
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
