package data

import (
	"lami/app/config"
	"lami/app/features/comments"
	event "lami/app/features/events/data"
	user "lami/app/features/users/data"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	db := config.InitDBTest()
	db.Migrator().DropTable(&Comment{}, &event.Event{}, &user.User{})
	db.AutoMigrate(&user.User{}, &event.Event{}, &Comment{})
	db.Create(&user.User{
		Name:   "Yusup",
		RoleID: 2,
	})
	db.Create(&event.Event{
		Name:   "Budaya",
		UserID: 1,
	})
	db.Create(&Comment{
		EventID: 1,
		UserID:  1,
		Comment: "Apakah masih bisa ?",
	})

	repo := NewCommentRepository(db)
	t.Run("Success Insert Comment", func(t *testing.T) {
		comment := comments.Core{
			EventID: 1,
			UserID:  1,
			Comment: "Apakah masih bisa ?",
		}

		err := repo.Insert(comment)
		assert.Nil(t, err)
	})

	t.Run("Failed Insert Comment No eventID", func(t *testing.T) {
		comment := comments.Core{
			EventID: 0,
			UserID:  1,
			Comment: "Apakah masih bisa ?",
		}

		err := repo.Insert(comment)
		assert.NotNil(t, err)
	})

	t.Run("Failed Insert Comment No userID", func(t *testing.T) {
		comment := comments.Core{
			EventID: 1,
			UserID:  0,
			Comment: "Apakah masih bisa ?",
		}

		err := repo.Insert(comment)
		assert.NotNil(t, err)
	})
}
