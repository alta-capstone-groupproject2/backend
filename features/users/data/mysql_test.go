package data

import (
	"lami/app/config"
	"lami/app/features/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectDataById(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})

	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})

	repo := NewUserRepository(db)
	t.Run("Success Get User", func(t *testing.T) {

		_, err := repo.SelectDataById(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Fail Get User", func(t *testing.T) {

		_, err := repo.SelectDataById(0)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestInsertData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})

	repo := NewUserRepository(db)
	t.Run("Success Insert User", func(t *testing.T) {
		user := users.Core{
			Name:     "alfin",
			Email:    "alfin2@mail.com",
			Password: "123",
			RoleID:   2,
		}
		err := repo.InsertData(user)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})

	t.Run("Fail Insert User Already Exist", func(t *testing.T) {
		user := users.Core{
			Name:     "alfin",
			Email:    "alfin@mail.com",
			Password: "123",
			RoleID:   2,
		}
		err := repo.InsertData(user)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

	t.Run("Fail Insert User", func(t *testing.T) {
		user := users.Core{
			Name:     "alfin",
			Email:    "al@mail.com",
			Password: "123",
		}
		err := repo.InsertData(user)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
}

func TestDeleteData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})
	repo := NewUserRepository(db)
	t.Run("Success Delete User", func(t *testing.T) {

		err := repo.DeleteData(1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Fail Delete User", func(t *testing.T) {

		err := repo.DeleteData(2)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})

}

func TestUpdateData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})
	repo := NewUserRepository(db)
	t.Run("Success Update User", func(t *testing.T) {
		data := map[string]interface{}{
			"name": "updated",
		}
		err := repo.UpdateData(data, 1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Fail Update User", func(t *testing.T) {
		data := map[string]interface{}{}
		err := repo.UpdateData(data, 2)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
}

func TestInsertStoreData(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})
	repo := NewUserRepository(db)

	t.Run("Success Inset User Store", func(t *testing.T) {
		data := users.Core{
			StoreName: "store A",
			Phone:     "91231231",
			Owner:     "my",
			City:      "malang",
			Address:   "simpang remujung",
			Document:  "url.example.com",
		}
		err := repo.InsertStoreData(data, 1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
}

func TestUpdateAccountRole(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:     "alfin",
		Email:    "alfin@mail.com",
		Password: "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:   2,
	})
	repo := NewUserRepository(db)
	t.Run("Success Update User Status", func(t *testing.T) {

		err := repo.UpdateAccountRole("approve", 2, 1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Fail Update User Status", func(t *testing.T) {

		err := repo.UpdateAccountRole("approve", 2, 2)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
}

func TestSelectDataSubmissionStore(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{
		Name:        "alfin",
		Email:       "alfin@mail.com",
		Password:    "$2a$10$NAwYGyMkMUK801hH84TdoeId9ZCZ6/0qO/Ao9a/mjg6CSCxSaKHFu",
		RoleID:      3,
		StoreStatus: "approve",
	})
	repo := NewUserRepository(db)
	t.Run("Success Get User Submission", func(t *testing.T) {

		_, _, err := repo.SelectDataSubmissionStore(1, 1)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
}
