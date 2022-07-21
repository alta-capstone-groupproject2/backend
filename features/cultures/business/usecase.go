package business

import (
	"errors"
	"fmt"
	"lami/app/config"
	"lami/app/features/cultures"
	"mime/multipart"
)

type cultureUseCase struct {
	cultureData cultures.Data
}

func NewCultureBusiness(prdData cultures.Data) cultures.Business {
	return &cultureUseCase{
		cultureData: prdData,
	}
}

// AddCulture implements culture.Business
func (uc *cultureUseCase) AddCulture(dataReq cultures.Core, fileInfo *multipart.FileHeader, fileData multipart.File) error {
	if dataReq.Name == "" || dataReq.City == "" || dataReq.Details == "" {
		return errors.New("all data must be filled")
	}

	if fileInfo != nil {
		urlImage, errFile := uploadFileValidation(dataReq.Name, 0, config.CultureImages, config.ContentImage, fileInfo, fileData)
		if errFile != nil {
			return errors.New(errFile.Error())
		}

		dataReq.Image = urlImage
	}

	err := uc.cultureData.AddDataCulture(dataReq)
	if err != nil {
		return errors.New("failed to insert data culture")
	}

	return nil
}

// SelectMyculture implements culture.Business
func (uc *cultureUseCase) SelectCulture(limit, page int) ([]cultures.Core, int64, error) {
	offset := limit * (page - 1)
	resp, total, err := uc.cultureData.SelectDataCulture(limit, offset)
	total = total/int64(limit) + 1
	return resp, total, err
}

// SelectculturebyCultureID implements culture.Business
func (uc *cultureUseCase) SelectCulturebyCultureID(cultureID int) (cultures.Core, error) {
	resp, err := uc.cultureData.SelectDataCultureByCultureID(cultureID)
	return resp, err
}

// DeleteCulture implements culture.Business
func (uc *cultureUseCase) DeleteCulture(cultureID int) error {
	err := uc.cultureData.DeleteDataCulture(cultureID)
	return err
}

// UpdateCulture implements culture.Business
func (uc *cultureUseCase) UpdateCulture(dataReq cultures.Core, cultureID int, fileInfo *multipart.FileHeader, fileData multipart.File) error {
	updateMap := make(map[string]interface{})

	if dataReq.Name != "" || dataReq.Name == " " {
		updateMap["name"] = &dataReq.Name
	}
	if dataReq.City != "" || dataReq.City == " " {
		updateMap["city"] = &dataReq.City
	}
	if dataReq.Details != "" || dataReq.Details == " " {
		updateMap["details"] = &dataReq.Details
	}
	if fileInfo != nil {
		urlImage, errFile := uploadFileValidation("", 0, config.CultureImages, config.ContentImage, fileInfo, fileData)
		if errFile != nil {
			return errors.New(errFile.Error())
		}

		updateMap["image"] = urlImage
	}
	fmt.Println(updateMap["image"])
	err := uc.cultureData.UpdateDataCulture(updateMap, cultureID)
	if err != nil {
		return errors.New("failed to update data culture")
	}

	return nil
}

// AddCultureReport implements culture.Business
func (uc *cultureUseCase) AddCultureReport(dataReq cultures.CoreReport) error {
	err := uc.cultureData.AddCultureDataReport(dataReq)
	if err != nil {
		return errors.New("failed to insert data report culture")
	}
	// userData, _ := uc.cultureData.SelectUser(dataReq.UserID)

	// helper.SendGmailNotify(userData.Email, "Success Add Report")
	return nil
}

// SelectReport implements culture.Business
func (uc *cultureUseCase) SelectReport(cultureID int) ([]cultures.CoreReport, error) {
	resp, err := uc.cultureData.SelectDataReport(cultureID)
	return resp, err
}
