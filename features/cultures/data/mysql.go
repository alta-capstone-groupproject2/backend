package data

import (
	"errors"
	"lami/app/features/cultures"
	"lami/app/features/users/data"

	"gorm.io/gorm"
)

type mysqlCultureRepository struct {
	db *gorm.DB
}

func NewCultureRepository(conn *gorm.DB) cultures.Data {
	return &mysqlCultureRepository{
		db: conn,
	}
}

// AddDataCulture implements culture.Data
func (repo *mysqlCultureRepository) AddDataCulture(dataReq cultures.Core) error {

	model := fromCore(dataReq)

	res := repo.db.Create(&model)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("failed insert data")
	}

	return nil

}

// SelectDataMyCulture implements culture.Data
func (repo *mysqlCultureRepository) SelectDataCulture(limit, offset int) ([]cultures.Core, int64, error) {
	dataCulture := []Culture{}
	var count int64
	res := repo.db.Order("updated_at desc").Limit(limit).Offset(offset).Find(&dataCulture).Count(&count)
	if res.Error != nil {
		return []cultures.Core{}, 0, errors.New("failed get data culture")
	}

	return ToCoreList(dataCulture), count, nil
}

// SelectDataCulturebyIDCulture implements culture.Data
func (repo *mysqlCultureRepository) SelectDataCultureByCultureID(idCulture int) (cultures.Core, error) {
	dataCulture := Culture{}

	res := repo.db.Where("id = ?", idCulture).Find(&dataCulture)
	if res.Error != nil {
		return cultures.Core{}, res.Error
	}

	return ToCore(dataCulture), res.Error
}

// DeleteDataCulture implements culture.Data
func (repo *mysqlCultureRepository) DeleteDataCulture(idCulture int) error {
	dataCulture := Culture{}

	res := repo.db.Delete(&dataCulture, idCulture)
	if res.RowsAffected == 0 {
		return errors.New("dailed delete culture")
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// UpdateDataCulture implements culture.Data
func (repo *mysqlCultureRepository) UpdateDataCulture(dataReq map[string]interface{}, idCulture int) error {

	model := Culture{}
	model.ID = uint(idCulture)
	err := repo.db.Model(model).Updates(dataReq)
	if err.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if err != nil {
		return err.Error
	}

	return nil
}

// AddCultureDataReport implements culture.Data
func (repo *mysqlCultureRepository) AddCultureDataReport(dataReq cultures.CoreReport) error {

	model := fromCoreReport(dataReq)

	res := repo.db.Create(&model)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("failed insert report")
	}

	return nil
}

// SelectDataReport implements culture.Data
func (repo *mysqlCultureRepository) SelectDataReport(idCulture int) ([]cultures.CoreReport, error) {
	dataReport := []Report{}

	res := repo.db.Preload("Culture").Where("culture_id = ?", idCulture).Find(&dataReport)
	if res.Error != nil {
		return []cultures.CoreReport{}, res.Error
	}

	return ToCoreReportList(dataReport), nil
}

func (repo *mysqlCultureRepository) SelectUser(id int) (response data.User, err error) {
	datauser := data.User{}
	result := repo.db.Preload("Role").Find(&datauser, id)
	if result.Error != nil {
		return data.User{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.User{}, errors.New("user not found")
	}
	return datauser, nil
}
