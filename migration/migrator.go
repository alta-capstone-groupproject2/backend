package migration

import (
	_mComment "lami/app/features/comments/data"
	_mEvent "lami/app/features/events/data"
	_mParticipant "lami/app/features/participants/data"
	_mUser "lami/app/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mEvent.Event{})
	db.AutoMigrate(_mParticipant.Participant{})
	db.AutoMigrate(_mComment.Comment{})
}
