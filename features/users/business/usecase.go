package business

import (
	"errors"
	"lami/app/features/users"

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

// func (uc *userUseCase) GetAllData(limit, offset int) (response []users.Core, err error) {
// 	resp, errData := uc.userData.SelectData(limit, offset)
// 	return resp, errData
// }

func (uc *userUseCase) GetDataById(id int) (response users.Core, err error) {
	resp, errData := uc.userData.SelectDataById(id)
	return resp, errData
}

func (uc *userUseCase) InsertData(userRequest users.Core) (row int, err error) {

	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		return -2, errors.New("all data must be filled")
	}

	passWillBcrypt := []byte(userRequest.Password)
	hash, err_hash := bcrypt.GenerateFromPassword(passWillBcrypt, bcrypt.DefaultCost)
	if err_hash != nil {
		return -1, errors.New("hashing password failed")
	}
	userRequest.Password = string(hash)
	result, err := uc.userData.InsertData(userRequest)
	if err != nil {
		return 0, errors.New("failed to insert data")
	}
	return result, nil
}

func (uc *userUseCase) DeleteData(id int) (row int, err error) {
	result, err := uc.userData.DeleteData(id)
	if err != nil {
		return 0, errors.New("no data user for deleted")
	}
	return result, nil
}

func (uc *userUseCase) UpdateData(userReq users.Core, id int) (row int, err error) {
	updateMap := make(map[string]interface{})

	if userReq.Name != "" {
		updateMap["name"] = &userReq.Name
	}

	if userReq.Email != "" {
		updateMap["email"] = &userReq.Email
	}

	if userReq.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
		if err != nil {
			return -1, errors.New("hasing password failed")
		}
		updateMap["password"] = &hash
	}

	if userReq.URL != "" {
		updateMap["url"] = &userReq.URL
	}

	result, err := uc.userData.UpdateData(updateMap, id)
	if err != nil {
		return 0, errors.New("no data user for updated")
	}
	return result, nil
}
