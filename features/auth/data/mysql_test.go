package data

import (
	"lami/app/config"
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

	repo := NewAuthRepository(db)
	t.Run("Success Get User", func(t *testing.T) {
		authData := User{
			Email:    "alfin@mail.com",
			Password: "123",
		}

		_, err := repo.FindUser(authData.Email)
		assert.Nil(t, err)
		//assert.Equal(t, 1, row)
	})
	t.Run("Success Get User", func(t *testing.T) {
		authData := User{
			Email:    "fail@mail.com",
			Password: "123",
		}

		_, err := repo.FindUser(authData.Email)
		assert.NotNil(t, err)
		//assert.Equal(t, 1, row)
	})
}
