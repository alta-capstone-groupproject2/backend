package data

import (
	"lami/app/config"
	"lami/app/features/cultures"
	"lami/app/features/users/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectDataCulture(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Culture{})
	db.AutoMigrate(&Culture{})
	db.Create(&Culture{
		Name: "culture A",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get User Submission", func(t *testing.T) {

		_, _, err := repo.SelectDataCulture(1, 1, "", "")
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})

	db.Migrator().DropTable(&Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get User Submission", func(t *testing.T) {

		_, _, err := repo.SelectDataCulture(1, 1, "", "")
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestSelectUser(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&data.User{})
	db.AutoMigrate(&data.User{})

	db.Create(&data.User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get User", func(t *testing.T) {

		_, err := repo.SelectUser(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Fail Get User", func(t *testing.T) {

		_, err := repo.SelectUser(0)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
	db.Migrator().DropTable(&data.User{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get User", func(t *testing.T) {

		_, err := repo.SelectUser(1)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestAddDataCulture(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Culture{})
	db.AutoMigrate(&Culture{})

	repo := NewCultureRepository(db)
	t.Run("Success Get User Submission", func(t *testing.T) {
		data := cultures.Core{
			ID:      0,
			Name:    "culture test",
			Image:   "url.com",
			City:    "malang",
			Details: "detail",
		}
		err := repo.AddDataCulture(data)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	db.Migrator().DropTable(&Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get User Submission", func(t *testing.T) {
		data := cultures.Core{
			ID:      0,
			Name:    "culture test",
			Image:   "url.com",
			City:    "malang",
			Details: "detail",
		}
		err := repo.AddDataCulture(data)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
}

func TestAddCultureDataReport(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Report{}, &Culture{})
	db.AutoMigrate(&Culture{}, &Report{})
	db.Create(&Culture{
		Name: "culture A",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get report", func(t *testing.T) {
		data := cultures.CoreReport{
			CultureID: 1,
			Message:   "report",
		}
		err := repo.AddCultureDataReport(data)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})

	db.Migrator().DropTable(&Report{}, &Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get Report", func(t *testing.T) {
		data := cultures.CoreReport{
			CultureID: 1,
			Message:   "report",
		}
		err := repo.AddCultureDataReport(data)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestSelectDataCultureByCultureID(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Culture{})
	db.AutoMigrate(&Culture{})

	db.Create(&Culture{
		Name: "Culture A",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get Culture detail", func(t *testing.T) {

		_, err := repo.SelectDataCultureByCultureID(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	db.Migrator().DropTable(&Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get Culture", func(t *testing.T) {

		_, err := repo.SelectDataCultureByCultureID(1)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestSelectDataCultureReport(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Report{}, &Culture{})
	db.AutoMigrate(&Culture{}, &Report{})
	db.Create(&Culture{
		Name: "culture A",
	})
	db.Create(&Report{
		CultureID: 1,
		Message:   "Report",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get User Submission", func(t *testing.T) {

		_, err := repo.SelectDataReport(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})

	db.Migrator().DropTable(&Report{}, &Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get User Submission", func(t *testing.T) {

		_, err := repo.SelectDataReport(2)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestDeleteDataCulture(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Culture{})
	db.AutoMigrate(&Culture{})

	db.Create(&Culture{
		Name: "Culture A",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Get Culture detail", func(t *testing.T) {

		err := repo.DeleteDataCulture(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	db.Migrator().DropTable(&Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get Culture", func(t *testing.T) {

		err := repo.DeleteDataCulture(1)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestUpdateDataCulture(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Culture{})
	db.AutoMigrate(&Culture{})

	db.Create(&Culture{
		Name: "Culture A",
		City: "malang",
	})

	repo := NewCultureRepository(db)
	t.Run("Success Update Culture", func(t *testing.T) {
		data := map[string]interface{}{
			"name": "culture Update",
		}
		err := repo.UpdateDataCulture(data, 1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	db.Migrator().DropTable(&Culture{})
	repo = NewCultureRepository(db)
	t.Run("Fail Get Culture", func(t *testing.T) {
		data := map[string]interface{}{
			"name": "culture Update",
		}
		err := repo.UpdateDataCulture(data, 1)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}
