package business

import (
	"errors"
	"lami/app/features/cultures"
	"lami/app/features/users/data"
	"lami/app/mocks"
	"mime/multipart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectCulture(t *testing.T) {
	repo := new(mocks.CultureData)
	returnData := []cultures.Core{{ID: 1, Name: "culture A"}}
	total := int64(1)
	limit := 10
	page := 1
	name := "bali"
	city := "bali"
	t.Run("Success Get All", func(t *testing.T) {
		repo.On("SelectDataCulture", 10, 0, "bali", "bali").Return(returnData, total, nil)

		srv := NewCultureBusiness(repo)

		res, _, err := srv.SelectCulture(limit, page, name, city)
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

}

func TestSelectCulturebyCultureID(t *testing.T) {
	repo := new(mocks.CultureData)
	returnData := cultures.Core{ID: 1, Name: "culture A"}
	id := 1
	t.Run("Success Get Culture Detail", func(t *testing.T) {
		repo.On("SelectDataCultureByCultureID", mock.Anything).Return(returnData, nil)

		srv := NewCultureBusiness(repo)

		res, err := srv.SelectCulturebyCultureID(id)
		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, res.ID)
		repo.AssertExpectations(t)
	})

}

func TestSelectReport(t *testing.T) {
	repo := new(mocks.CultureData)
	returnData := []cultures.CoreReport{{ID: 1, Message: "Report A"}}
	id := 1
	t.Run("Success Get Culture Report", func(t *testing.T) {
		repo.On("SelectDataReport", mock.Anything).Return(returnData, nil)

		srv := NewCultureBusiness(repo)

		res, err := srv.SelectReport(id)
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

}

func TestDeleteCulture(t *testing.T) {
	repo := new(mocks.CultureData)
	id := 1
	t.Run("Success Delete Culture", func(t *testing.T) {
		repo.On("DeleteDataCulture", mock.Anything).Return(nil)

		srv := NewCultureBusiness(repo)

		err := srv.DeleteCulture(id)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}

func TestAddCultureReport(t *testing.T) {
	repo := new(mocks.CultureData)
	id := 1
	reportData := cultures.CoreReport{
		CultureID: id,
		UserID:    id,
		Message:   "report",
	}

	t.Run("Success Add Culture Report", func(t *testing.T) {
		repo.On("AddCultureDataReport", reportData).Return(nil).Once()
		repo.On("SelectDataCultureByCultureID", 1).Return(cultures.Core{}, nil).Once()
		repo.On("SelectUser", 1).Return(data.User{}, nil).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCultureReport(reportData)

		assert.NoError(t, errAdd)
		repo.AssertExpectations(t)
	})
	t.Run("Fail Add Culture Report", func(t *testing.T) {
		repo.On("AddCultureDataReport", reportData).Return(errors.New("report failed")).Once()
		repo.On("SelectDataCultureByCultureID", 1).Return(cultures.Core{}, nil).Once()
		repo.On("SelectUser", 1).Return(data.User{}, nil).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCultureReport(reportData)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})
}

func TestAddCulture(t *testing.T) {
	repo := new(mocks.CultureData)

	file, _ := os.Open("./../../../helper/testFile/imageSuccess.png")
	fileStat, _ := file.Stat()
	fileInfo := &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size:     fileStat.Size(),
	}
	fileData, _ := fileInfo.Open()

	fileSize, _ := os.Open("./../../../helper/testFile/imageSizeInvalid.png")
	fileStatSize, _ := fileSize.Stat()
	fileInfoSize := &multipart.FileHeader{
		Filename: fileStatSize.Name(),
		Size:     fileStatSize.Size(),
	}

	fileDataSize, _ := fileInfoSize.Open()
	fileFail, _ := os.Open("./../../../helper/testFile/fileTest.txt")
	fileStatFail, _ := fileFail.Stat()
	fileInfoFail := &multipart.FileHeader{
		Filename: fileStatFail.Name(),
		Size:     fileStatFail.Size(),
	}
	fileDataFail, _ := fileInfoFail.Open()

	id := 1
	data := cultures.Core{
		ID:      id,
		Name:    "jaran kepang",
		Image:   "urlImage",
		City:    "malang",
		Details: "detail jaran kepang",
	}

	t.Run("Success Add Culture Report", func(t *testing.T) {
		repo.On("AddDataCulture", data).Return(nil).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCulture(data, fileInfo, fileData)

		assert.NoError(t, errAdd)
		repo.AssertExpectations(t)
	})
	t.Run("Fail Add Culture size invalid", func(t *testing.T) {
		repo.On("AddDataCulture", data).Return(errors.New("add failed")).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCulture(data, fileInfoSize, fileDataSize)

		assert.NotNil(t, errAdd)
	})
	t.Run("Fail Add Culture Report", func(t *testing.T) {
		repo.On("AddDataCulture", data).Return(errors.New("add failed")).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCulture(data, fileInfoFail, fileDataFail)

		assert.NotNil(t, errAdd)
	})
	t.Run("Fail Add Culture", func(t *testing.T) {
		repo.On("AddDataCulture", data).Return(errors.New("add failed")).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.AddCulture(data, fileInfo, fileData)

		assert.NotNil(t, errAdd)
	})
	t.Run("Fail Add Culture empty", func(t *testing.T) {
		repo.On("AddDataCulture", data).Return(errors.New("add failed")).Once()

		srv := NewCultureBusiness(repo)
		data.City = ""
		errAdd := srv.AddCulture(data, fileInfo, fileData)

		assert.NotNil(t, errAdd)
	})

}

func TestUpdateCulture(t *testing.T) {
	repo := new(mocks.CultureData)

	file, _ := os.Open("./../../../helper/testFile/imageSuccess.png")
	fileStat, _ := file.Stat()
	fileInfo := &multipart.FileHeader{
		Filename: fileStat.Name(),
		Size:     fileStat.Size(),
	}
	fileData, _ := fileInfo.Open()

	fileSize, _ := os.Open("./../../../helper/testFile/imageSizeInvalid.png")
	fileStatSize, _ := fileSize.Stat()
	fileInfoSize := &multipart.FileHeader{
		Filename: fileStatSize.Name(),
		Size:     fileStatSize.Size(),
	}

	fileDataSize, _ := fileInfoSize.Open()
	fileFail, _ := os.Open("./../../../helper/testFile/fileTest.txt")
	fileStatFail, _ := fileFail.Stat()
	fileInfoFail := &multipart.FileHeader{
		Filename: fileStatFail.Name(),
		Size:     fileStatFail.Size(),
	}
	fileDataFail, _ := fileInfoFail.Open()

	id := 1
	data := cultures.Core{
		ID:      id,
		Name:    "jaran kepang",
		Image:   "urlImage",
		City:    "malang",
		Details: "detail jaran kepang",
	}

	t.Run("Success update Culture", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return(nil).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.UpdateCulture(data, id, fileInfo, fileData)

		assert.NoError(t, errAdd)
		// repo.AssertExpectations(t)
	})
	t.Run("Fail update Culture", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return(errors.New("failed update")).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.UpdateCulture(data, id, fileInfo, fileData)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})
	t.Run("Fail Update Culture Size invalid", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return("invalid file").Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.UpdateCulture(data, id, fileInfoSize, fileDataSize)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})
	t.Run("Fail Update Culture Extension invalid", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return(errors.New("invalid file")).Once()

		srv := NewCultureBusiness(repo)

		errAdd := srv.UpdateCulture(data, id, fileInfoFail, fileDataFail)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})

	t.Run("Fail update Culture city", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return("invalid city").Once()

		srv := NewCultureBusiness(repo)
		data.City = "1"
		errAdd := srv.UpdateCulture(data, id, fileInfo, fileData)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})
	t.Run("Fail update Culture name", func(t *testing.T) {
		repo.On("UpdateDataCulture", mock.Anything, 1).Return("invalid name").Once()

		srv := NewCultureBusiness(repo)
		data.City = "malang"
		data.Name = "1"
		errAdd := srv.UpdateCulture(data, id, fileInfo, fileData)

		assert.NotNil(t, errAdd)
		// repo.AssertExpectations(t)
	})

}
