package business

import (
	"errors"
	"lami/app/features/auth"
	"lami/app/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	userData auth.Data
}

func NewAuthBusiness(usrData auth.Data) auth.Business {
	return &authUseCase{
		userData: usrData,
	}
}

func (uc *authUseCase) Login(data auth.Core) (string, int, string, error) {
	response, errFind := uc.userData.FindUser(data.Email)
	if errFind != nil {
		return "", 0, "", errors.New("user not found")
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(data.Password))
	if errCompare != nil {
		return "", 0, "", errors.New("wrong password")
	}
	token, err := middlewares.CreateToken(int(response.ID), response.Role)

	return token, response.ID, response.Role, err
}
