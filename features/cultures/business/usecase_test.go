package business

import (
	"lami/app/features/cultures"
	"lami/app/features/users"
	"lami/app/mocks"
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
		repo.On("SelectDataCulture", 10,0, "bali", "bali").Return(returnData, total, nil)

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

	t.Run("Success Add Culture Report", func(t *testing.T) {
		repo.On("AddCultureReport", cultures.CoreReport{
			CultureID: 1,
			UserID: 1,
			Message: "report",
		}).Return(nil).Once()
		repo.On("SelectDataCultureByCultureID", 1).Return(cultures.Core{}).Once()
		repo.On("SelectUser", 1).Return(users.Core{}).Once()
		reportData := cultures.CoreReport{
			CultureID: id,
			UserID:    id,
			Message:   "report",
		}

		srv := NewCultureBusiness(repo)

		err := srv.AddCultureReport(reportData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}


