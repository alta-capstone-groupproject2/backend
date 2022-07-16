package data

import (
	"errors"
	"lami/app/features/cultures"

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

// SelectDataReport implements culture.Data
func (repo *mysqlCultureRepository) SelectDataReport(idCulture int) ([]cultures.CoreReport, error) {
	dataReport := []Report{}

	res := repo.db.Preload("Culture").Where("culture_id = ?", idCulture).Find(&dataReport)
	if res.Error != nil {
		return []cultures.CoreReport{}, res.Error
	}

	return ToCoreReport(dataReport), nil
}

// SelectDataMyCulture implements culture.Data
func (repo *mysqlCultureRepository) SelectDataMyCulture(idUser int) ([]cultures.Core, error) {
	dataMyCulture := []Culture{}

	res := repo.db.Preload("User").Where("user_id = ?", idUser).Find(&dataMyCulture)
	if res.Error != nil {
		return []cultures.Core{}, res.Error
	}

	return ToCoreMyCulture(dataMyCulture), nil
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
func (repo *mysqlCultureRepository) DeleteDataCulture(idCulture int, idUser int) error {
	dataCulture := Culture{}
	res := repo.db.Where("user_id = ?", idUser).Delete(&dataCulture, idCulture)
	if res.RowsAffected == 0 {
		return errors.New("no rows affected")
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

// UpdateDataCulture implements culture.Data
func (repo *mysqlCultureRepository) UpdateDataCulture(dataReq cultures.Core, idCulture int) error {

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
