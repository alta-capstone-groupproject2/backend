package business

import (
	"errors"
	"lami/app/config"
	"lami/app/features/users"
	"lami/app/helper"
	"mime/multipart"
	"strings"

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
	hash, _ := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	
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
		hash, _ := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		
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

	errOwner := nameFormatValidation(dataReq.Owner)
	if errOwner != nil {
		return errors.New(errOwner.Error())
	}

	errCity := cityFormatValidation(dataReq.City)
	if errCity != nil {
		return errors.New(errCity.Error())
	}

	errPhone := phoneFormatValidation(dataReq.Phone)
	if errPhone != nil {
		return errors.New(errPhone.Error())
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

	//helper.SendGmailNotify(dataReq.Email, "upgrade account to umkm")
	return nil
}
func (uc *userUseCase) UpdateStatusUser(status string, id int) error {
	roleId := 0
	if status == "approve" {
		roleId = 3
	} else {
		roleId = 2
	}
	err := uc.userData.UpdateAccountRole(status, roleId, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) GetDataSubmissionStore(limit, page int) (response []users.Core, total int64, err error) {
	offset := limit * (page - 1)
	resp, total, errData := uc.userData.SelectDataSubmissionStore(limit, offset)
	total = total/int64(limit) + 1
	return resp, total, errData
}

func (uc *userUseCase) VerifyEmail(userData users.Core) error {
	//random string for sparator
	key := helper.RandomString(3)

	//combine data and sparator
	if userData.Name == "" || userData.Email == "" || userData.Password == "" || userData.Name == " " || userData.Password == " " {
		return errors.New("all data must be filled")
	}

	errEmailFormat := emailFormatValidation(userData.Email)
	if errEmailFormat != nil {
		return errors.New(errEmailFormat.Error())
	}

	errNameFormat := nameFormatValidation(userData.Name)
	if errNameFormat != nil {
		return errors.New(errNameFormat.Error())
	}

	plain := key + key + userData.Name + key + userData.Email + key + userData.Password

	encrypted := helper.Encrypt(plain, "57a45acad2047e9731ed4bd06c4f2af2f556d60da076606dea4d55463fdff03f")

	helper.SendEmailVerification(userData, encrypted)
	return nil
}

func (uc *userUseCase) ConfirmEmail(encryptData string) error {
	var userData users.Core
	Decrypted := helper.Decrypt(encryptData, "57a45acad2047e9731ed4bd06c4f2af2f556d60da076606dea4d55463fdff03f")

	// get sparator
	sparator := Decrypted[:6]
	dataRaw := strings.Split(Decrypted, sparator)
	userData.Name = dataRaw[2]
	userData.Email = dataRaw[3]
	userData.Password = dataRaw[4]

	userData.RoleID = 2
	userData.Image = "https://lamiapp.s3.amazonaws.com/userimages/default_user.png"
	err := uc.userData.InsertData(userData)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
