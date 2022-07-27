package business

import (
	"fmt"
	"lami/app/features/users"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//mock data success case
type mockUserData struct{}

func (mock mockUserData) SelectDataById(id int) (data users.Core, err error) {
	return users.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}, nil
}

func (mock mockUserData) SelectDataSubmissionStore(limit int, offset int) (data []users.Core, total int64, err error) {
	return []users.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty"}}, 1,nil
}

func (mock mockUserData) InsertData(data users.Core) (err error) {
	return nil
}

func (mock mockUserData) InsertStoreData(data users.Core, id int) (err error) {
	return nil
}

func (mock mockUserData) DeleteData(id int) (err error) {
	return nil
}

func (mock mockUserData) UpdateData(data map[string]interface{}, id int) (err error) {
	return nil
}
func (mock mockUserData) UpdateAccountRole(status string, roledId, id int) (err error) {
	return nil
}

//mock data failed case
type mockUserDataFailed struct{}

func (mock mockUserDataFailed) SelectDataById(id int) (data users.Core, err error) {
	return data, fmt.Errorf("Failed to select data")
}

func (mock mockUserDataFailed) SelectDataSubmissionStore(limit int, offset int) (data []users.Core, total int64, err error) {
	return data, 1, fmt.Errorf("failed get submission data")
}

func (mock mockUserDataFailed) InsertData(data users.Core) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockUserDataFailed) DeleteData(id int) (err error) {
	return fmt.Errorf("failed to delete data ")
}

func (mock mockUserDataFailed) UpdateData(data map[string]interface{}, id int) (err error) {
	return fmt.Errorf("failed to update data ")
}

func (mock mockUserDataFailed) InsertStoreData(data users.Core, id int) (err error) {
	return fmt.Errorf("failed to insert data ")
}

func (mock mockUserDataFailed) UpdateAccountRole(status string, roledId, id int) (err error) {
	return fmt.Errorf("failed to delete data ")
}

func TestGetDataById(t *testing.T) {
	t.Run("Success Get Data User", func(t *testing.T) {
		id := 1
		userBusiness := NewUserBusiness(mockUserData{})
		result, err := userBusiness.GetDataById(id)
		assert.Nil(t, err)
		assert.Equal(t, "alta", result.Name)
	})
	t.Run("Failed Get Data User", func(t *testing.T) {
		id := 3
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result, err := userBusiness.GetDataById(id)
		assert.NotNil(t, err)
		assert.Equal(t, "", result.Name)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("Success Test Insert Data", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.Nil(t, err)
		
	})

	t.Run("Success Test Insert Data", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
		
	})

	t.Run("Failed Test Insert Data When Email Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email: "alta",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Insert Data When Email Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta1",
			Email: "alta@mail.com",
			Password: "qwerty",
		}
		err := userBusiness.InsertData(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Insert Data When All Data Not Filled ", func(t *testing.T) {
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
	t.Run("Success Test Delete Data", func(t *testing.T) {
		id := 1
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.DeleteData(id)
		assert.Nil(t, err)
		
	})
	t.Run("Failed Test Delete Data", func(t *testing.T) {
		id := 1
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.DeleteData(id)
		assert.NotNil(t, err)
		
	})
}

func TestUpdateData(t *testing.T) {
	file , _ := os.Open("./../../../helper/testFile/imageSuccess.png")
	fileStat, _ := file.Stat()
	fileInfo := &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size: fileStat.Size(),
		
	}
	fileData, _ := fileInfo.Open()
	
	fileSize , _ := os.Open("./../../../helper/testFile/imageSizeInvalid.png")
	fileStatSize, _ := fileSize.Stat()
	fileInfoSize := &multipart.FileHeader{
		Filename: fileStatSize.Name(),
		Size: fileStatSize.Size(),
		
	}

	fileDataSize, _ := fileInfoSize.Open()
	fileFail, _ := os.Open("./../../../helper/testFile/fileTest.txt")
		fileStatFail,_ :=fileFail.Stat()
		fileInfoFail := &multipart.FileHeader{
			Filename: fileStatFail.Name(),
			Size: fileStatFail.Size(),			
		}	
	fileDataFail, _ := fileInfoFail.Open()

	t.Run("Success Test Update Data", func(t *testing.T) {
		id := 1
		data := users.Core{
			Name:     "Alfin",
			Image:    "https://image.site",
			Email:    "alfin@mail.id",
			Password: "123",
		}
		
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.UpdateData(data, id, fileInfo, fileData)
		assert.Nil(t, err)
		
	})
	t.Run("Failed Test Update Name Invalid", func(t *testing.T) {
		id := 3
		data := users.Core{
			Name:     "Alfin1",
			Image:    "https://image.site",
			Email:    "alfin@mail.id",
			Password: "123",
		}
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateData(data, id, fileInfo, fileData)
		assert.NotNil(t, err)
		
	})
	t.Run("Failed Test Update Email Invalid", func(t *testing.T) {
		id := 3
		data := users.Core{
			Name:     "Alfin",
			Image:    "https://image.site",
			Email:    "alfin.id",
			Password: "123",
		}
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateData(data, id, fileInfo, fileData)
		assert.NotNil(t, err)
		
	})
	t.Run("Failed Test Update Data", func(t *testing.T) {
		id := 3
		data := users.Core{
			Name:     "Alfin",
			Image:    "https://image.site",
			Email:    "alfin@mail.id",
			Password: "123",
		}
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateData(data, id, fileInfo, fileData)
		assert.NotNil(t, err)	
	})
	t.Run("Failed Test Update Data invalid extension", func(t *testing.T) {
		id := 1
		data := users.Core{
			Name:     "Alfin",
			Image:    "https://image.site",
			Email:    "alfin@mail.id",
			Password: "123",
		}
		
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateData(data, id, fileInfoFail, fileDataFail)
		assert.NotNil(t, err)
		
	})
	t.Run("Failed Test Update Data invalid size", func(t *testing.T) {
		id := 1
		data := users.Core{
			Name:     "Alfin",
			Image:    "https://image.site",
			Email:    "alfin@mail.id",
			Password: "123",
		}
		
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateData(data, id, fileInfoSize, fileDataSize)
		assert.NotNil(t, err)
		
	})
	
	
}

func TestUpgradeAccount(t *testing.T) {
	file , _ := os.Open("./../../../helper/testFile/docSuccess.pdf")
	fileStat, _ := file.Stat()
	fileInfo := &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size: fileStat.Size(),
		
	}
	fileData, _ := fileInfo.Open()
	
	fileSize , _ := os.Open("./../../../helper/testFile/docInvalidSize.pdf")
	fileStatSize, _ := fileSize.Stat()
	fileInfoSize := &multipart.FileHeader{
		Filename: fileStatSize.Name(),
		Size: fileStatSize.Size(),
		
	}
	fileFail, _ := os.Open("./../../../helper/testFile/imageSuccess.png")
		fileStatFail,_ :=fileFail.Stat()
		fileInfoFail := &multipart.FileHeader{
			Filename: fileStatFail.Name(),
			Size: fileStatFail.Size(),			
		}	
	fileDataFail, _ := fileInfoFail.Open()

	
	t.Run("Success Test Upgrade Account", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku",
			City:        "malang",
			Address:     "jln. simpang remujung",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfo, fileData)
		assert.Nil(t, err)
		
	})

	t.Run("Failed Test Upgrade Account Not All Data Filled", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku",
			City:        "malang",
			Address:     "",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfo, fileData)
		assert.NotNil(t, err)
		
	})

	t.Run("Failed Test Upgrade Account When Name Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku1",
			City:        "malang",
			Address:     "jln. simpang remujung",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfo, fileData)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Upgrade Account When City Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku",
			City:        "malang1",
			Address:     "jln. simpang remujung",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfo, fileData)
		assert.NotNil(t, err)
	})
	t.Run("Failed Test Upgrade Account When All Data Not Filled ", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
		StoreName:   "store A",
		Phone:       "8218198212qq",
		Owner:       "aku",
		City:        "malang",
		Address:     "jln. simpang remujung",
	}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfo, fileData)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Upgrade Account When extension invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku",
			City:        "malang",
			Address:     "jln. simpang remujung",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfoFail, fileDataFail)
		assert.NotNil(t, err)
	})
	t.Run("Failed Test Upgrade Account When size invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newStoreUser := users.Core{
			StoreName:   "store A",
			Phone:       "8218198212",
			Owner:       "aku",
			City:        "malang",
			Address:     "jln. simpang remujung",
		}
		err := userBusiness.UpgradeAccount(newStoreUser, 1, fileInfoSize, fileData)
		assert.NotNil(t, err)
	})
}

func TestUpdateStatusUser(t *testing.T) {
	t.Run("Success Test Update Status Approve Data", func(t *testing.T) {
		id := 1
		status := "approve"
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.UpdateStatusUser(status,id)
		assert.Nil(t, err)
		
	})
	t.Run("Success Test Update Status Approve Data", func(t *testing.T) {
		id := 1
		status := "decline"
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.UpdateStatusUser(status,id)
		assert.Nil(t, err)
		
	})
	t.Run("Failed Test Delete Data", func(t *testing.T) {
		id := 0
		status := "decline"
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.UpdateStatusUser(status,id)
		assert.NotNil(t, err)
		
	})
}
func TestGetDataSubmissionStore(t *testing.T) {
	t.Run("Success Get Data Submission Store User", func(t *testing.T) {
		limit := 10
		offset := 1
		userBusiness := NewUserBusiness(mockUserData{})
		result, _, err := userBusiness.GetDataSubmissionStore(limit, offset)
		assert.Nil(t, err)
		// assert.NotEqual(t, total, int64(0))
		assert.Equal(t, "alta", result[0].Name)
	})
	t.Run("Failed Get Data Data Submission Store User", func(t *testing.T) {
		limit := 1
		offset := 1
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		result,_, err := userBusiness.GetDataSubmissionStore(limit, offset)
		assert.NotNil(t, err)
		// assert.Equal(t, total, int64(0))
		assert.Equal(t, []users.Core(nil), result)
	})
}

func TestVerifyEmail(t *testing.T) {
	t.Run("Success Test Verify", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserData{})
		newUser := users.Core{
			Name:     "alta",
			Email:    "alta@mail.id",
			Password: "qwerty",
		}
		err := userBusiness.VerifyEmail(newUser)
		assert.Nil(t, err)
		
	})


	t.Run("Failed Test Verify Email When Email Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta",
			Email: "alta",
			Password: "qwerty",
		}
		err := userBusiness.VerifyEmail(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Verify Email When Email Invalid", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:     "alta1",
			Email: "alta@mail.com",
			Password: "qwerty",
		}
		err := userBusiness.VerifyEmail(newUser)
		assert.NotNil(t, err)
	})

	t.Run("Failed Test Verify Email When All Data Not Filled ", func(t *testing.T) {
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		newUser := users.Core{
			Name:  "alta",
			Email: "alta@mail.id",
		}
		err := userBusiness.VerifyEmail(newUser)
		assert.NotNil(t, err)
	})
}

func TestConfirmEmail(t *testing.T) {
	t.Run("Success Test Delete Data", func(t *testing.T) {
		encrypt := "d86a98f985e4c6388996bf1a8e6322382d6b6f4d136cdaea78585f7c5678529c3eed34015d59cf4c6b3ce86711fe6d63598a346a2ae78b468f6fff8fb87926d33d890cc7e97a3e47ac6fe7f3dee10b"
		userBusiness := NewUserBusiness(mockUserData{})
		err := userBusiness.ConfirmEmail(encrypt)
		assert.Nil(t, err)
		
	})
	t.Run("Failed Test Delete Data", func(t *testing.T) {
		encrypt := "d86a98f985e4c6388996bf1a8e6322382d6b6f4d136cdaea78585f7c5678529c3eed34015d59cf4c6b3ce86711fe6d63598a346a2ae78b468f6fff8fb87926d33d890cc7e97a3e47ac6fe7f3dee10b"
		userBusiness := NewUserBusiness(mockUserDataFailed{})
		err := userBusiness.ConfirmEmail(encrypt)
		assert.NotNil(t, err)
		
	})
}