package business

// import (
// 	"errors"
// 	"lami/app/features/auth"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// type mockAuthDataSucsess struct{}

// func (mock mockAuthDataSucsess) FindUser(email string) (resp auth.Core, err error) {
// 	return auth.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "$2a$10$36DHlayBJtBvYZfGhP/bMOuU3zX9O7sz4ohSA1a6ikQIMVz05fobW"}, nil
// }

// type mockAuthDataFailed struct{}

// func (mock mockAuthDataFailed) FindUser(email string) (resp auth.Core, err error) {
// 	return auth.Core{}, errors.New("user not found")
// }

// func TestLogin(t *testing.T) {
// 	t.Run("Test Login Success", func(t *testing.T) {
// 		authBusiness := NewAuthBusiness(mockAuthDataSucsess{})
// 		newUser := auth.Core{
// 			ID:       1,
// 			Name:     "alta",
// 			Email:    "alta@mail.id",
// 			Password: "123",
// 		}
// 		resultToken, resultID, err := authBusiness.Login(newUser)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, resultToken)
// 		assert.Equal(t, 1, resultID)
// 	})

// 	t.Run("Test Login email not found", func(t *testing.T) {
// 		authBusiness := NewAuthBusiness(mockAuthDataFailed{})
// 		newUser := auth.Core{
// 			ID:       1,
// 			Name:     "alta",
// 			Email:    "abc@mail.id",
// 			Password: "123",
// 		}
// 		resultToken, resultID, err := authBusiness.Login(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", resultToken)
// 		assert.Equal(t, 0, resultID)
// 	})

// 	t.Run("Test Login Wrong Pass", func(t *testing.T) {
// 		authBusiness := NewAuthBusiness(mockAuthDataFailed{})
// 		newUser := auth.Core{
// 			ID:       1,
// 			Name:     "alta",
// 			Email:    "alta@mail.id",
// 			Password: "qwert1",
// 		}
// 		resultToken, result, err := authBusiness.Login(newUser)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, "", resultToken)
// 		if newUser.Password != "qwerty" {
// 			assert.Equal(t, 0, result)
// 		}
// 	})
// }
