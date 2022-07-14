package business

import (
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

func (uc *authUseCase) Login(data auth.Core) (string, int, error) {
	response, errFind := uc.userData.FindUser(data.Email)
	if errFind != nil {
		return "", 0, errFind
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(data.Password))
	if errCompare != nil {
		return "", 0, errCompare
	}
	token, err := middlewares.CreateToken(int(response.ID))

	return token, response.ID, err
}
