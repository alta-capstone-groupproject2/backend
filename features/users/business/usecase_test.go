package business

import (
	"fmt"
	"lami/app/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockUserData struct{}

/**
type Core struct {
    ID          int
    Name        string
    Email       string
    Password    string
    Image       string
    StoreName   string
    Phone       string
    Owner       string
    City        string
    Address     string
    Document    string
    RoleID      int
    StoreStatus string
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Role        Role
}
**/

func (mock mockUserData) SelectDataSubmissionStore(limit, offset int) (data []users.Core, total int64, err error) {
	return []users.Core{}, total, nil
}

func (mock mockUserData) SelectDataById(id int) (data users.Core, err error) {
	return users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}, nil
}

func (mock mockUserData) InsertData(data users.Core) error {
	return nil
}

func (mock mockUserData) DeleteData(param int) error {
	return nil
}

func (mock mockUserData) UpdateData(data map[string]interface{}, id int) error {
	return nil
}

func (mock mockUserData) InsertStoreData(dataReq users.Core, id int) error {
	return nil
}

func (mock mockUserData) UpdateAccountRole(string, int, int) error {
	return nil
}

//mock data failed case
type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectDataSubmissionStore(limit, offset int) (data []users.Core, total int64, err error) {
	return nil, 0, fmt.Errorf("failed get data submission")
}

func (mock mockUserDataFailed) SelectDataById(id int) (data users.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) InsertData(data users.Core) error {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockUserDataFailed) DeleteData(id int) error {
	return fmt.Errorf("failed to delete data ")
}

func (mock mockUserDataFailed) UpdateData(data map[string]interface{}, id int) error {
	return fmt.Errorf("failed to update data ")
}

func (mock mockUserDataFailed) InsertStoreData(dataReq users.Core, id int) error {
	return fmt.Errorf("failed insert data")
}

func (mock mockUserDataFailed) UpdateAccountRole(string, int, int) error {
	return fmt.Errorf("failed update data")
}

func TestGetDataById(t *testing.T) {
	t.Run("Test Get Data By Id", func(t *testing.T) {
		id := 1
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetDataById(id)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result.Name)
	})
	t.Run("Test Get Data By Id Failed", func(t *testing.T) {
		id := 3
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetDataById(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("Test Insert Data Success", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.Nil(t, err)
	})

	t.Run("Test Insert Data Failed", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Test Insert Data Failed When Email Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Test Insert Data Failed When Password Empty", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:  "alta",
			Email: "alta@mail.id",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
	})
}

func TestDeleteData(t *testing.T) {
	t.Run("Test Delete Data", func(t *testing.T) {
		param := 1
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.DeleteData(param)
		assert.Nil(t, err)
	})
	t.Run("Test Delete Data Failed", func(t *testing.T) {
		id := 1
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.DeleteData(id)
		assert.NotNil(t, err)
	})
}

// func TestUpdateData(t *testing.T) {
// 	t.Run("Test Update Data", func(t *testing.T) {
// 		id := 1
// 		data := users.Core{
// 			Name:     "Zaki",
// 			Image:    "https://image.site",
// 			Email:    "mail@mail.id",
// 			Password: "Mail@",
// 		}
// 		userBusiness := NewUserBusiness(mockUserData{})
// 		err := userBusiness.UpdateData(data, id)
// 		assert.Nil(t, err)
// 	})
// 	t.Run("Test Update Data Failed No data", func(t *testing.T) {
// 		id := 3
// 		data := users.Core{
// 			Name:     "Zaki",
// 			Image:    "https://image.site",
// 			Email:    "mail@mail.id",
// 			Password: "Mail@",
// 		}
// 		userBusiness := NewUserBusiness(mockUserDataFailed{})
// 		err := userBusiness.UpdateData(data, id)
// 		assert.NotNil(t, err)
// 	})
// }
