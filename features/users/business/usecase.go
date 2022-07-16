package business

import (
	"errors"
	"lami/app/config"
	"lami/app/features/users"
	"mime/multipart"

	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(usrData users.Data) users.Business {
	return &userUseCase{
		userData: usrData,
	}
}

func (uc *userUseCase) GetDataById(id int) (response users.Core, err error) {
	resp, errData := uc.userData.SelectDataById(id)
	return resp, errData
}

func (uc *userUseCase) InsertData(userRequest users.Core) (err error) {

	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" || userRequest.Name == " " || userRequest.Password == " " {
		return errors.New("all data must be filled")
	}

	errEmailFormat := emailFormatValidation(userRequest.Email)
	if errEmailFormat != nil {
		return errors.New(errEmailFormat.Error())
	}

	errNameFormat := nameFormatValidation(userRequest.Name)
	if errNameFormat != nil {
		return errors.New(errNameFormat.Error())
	}

	passWillBcrypt := []byte(userRequest.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return errors.New("hashing password failed")
	}
	userRequest.Password = string(hash)

	//default role user
	userRequest.RoleID = 2
	userRequest.Image = "https://lamiapp.s3.amazonaws.com/userimages/default_user.png"
	err = uc.userData.InsertData(userRequest)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (uc *userUseCase) DeleteData(id int) (err error) {
	err = uc.userData.DeleteData(id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) UpdateData(userReq users.Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) (err error) {
	updateMap := make(map[string]interface{})

	if userReq.Name != "" || userReq.Name == " " {
		errNameFormat := nameFormatValidation(userReq.Name)
		if errNameFormat != nil {
			return errors.New(errNameFormat.Error())
		}
		updateMap["name"] = &userReq.Name
	}
	if userReq.Email != "" {
		errEmailFormat := emailFormatValidation(userReq.Email)
		if errEmailFormat != nil {
			return errors.New(errEmailFormat.Error())
		}
		updateMap["email"] = &userReq.Email
	}
	if userReq.Password != "" {
		hash, errHash := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if errHash != nil {
			return errors.New("hasing password failed")
		}
		updateMap["password"] = &hash
	}

	if fileInfo != nil {
		urlImage, errFile := uploadFileValidation(userReq.Name, id, config.UserImages, config.ContentImage, fileInfo, fileData)
		if errFile != nil {
			return errors.New(errFile.Error())
		}

		updateMap["image"] = urlImage
	}

	err = uc.userData.UpdateData(updateMap, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) UpgradeAccount(dataReq users.Core, id int, fileInfo *multipart.FileHeader, fileData multipart.File) error {
	if dataReq.StoreName == "" || dataReq.Phone == "" || dataReq.Owner == "" || dataReq.City == "" || dataReq.Address == "" || fileInfo == nil {
		return errors.New("all data must be filled")
	}

	urlDoc, errFile := uploadFileValidation(dataReq.StoreName, id, config.UserDocuments, config.ContentDocuments, fileInfo, fileData)
	if errFile != nil {
		return errors.New(errFile.Error())
	}
	dataReq.Document = urlDoc
	dataReq.StoreStatus = "waiting"

	err := uc.userData.InsertStoreData(dataReq, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) UpdateStatusUser(status string, id int) error {
	err := uc.userData.UpdateAccountRole(status, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) GetAllData(limit, page int) (response []users.Core, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.userData.SelectData(limit, offset)
	total = total/int64(limit) + 1
	return resp, total, errData
}
