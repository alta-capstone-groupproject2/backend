package migration

import (
	_mCart "lami/app/features/carts/data"
	_mComment "lami/app/features/comments/data"
	_mCulture "lami/app/features/cultures/data"
	_mEvent "lami/app/features/events/data"
	_mOrder "lami/app/features/orders/data"
	_mParticipant "lami/app/features/participants/data"
	_mProduct "lami/app/features/products/data"
	_mUser "lami/app/features/users/data"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(_mUser.User{})
	db.AutoMigrate(_mUser.Role{})
	db.AutoMigrate(_mEvent.Event{})
	db.AutoMigrate(_mParticipant.Participant{})
	db.AutoMigrate(_mComment.Comment{})
	db.AutoMigrate(_mCulture.Culture{})
	db.AutoMigrate(_mCulture.Report{})
	db.AutoMigrate(_mProduct.Product{})
	db.AutoMigrate(_mProduct.Rating{})
	db.AutoMigrate(_mCart.Cart{})
	db.AutoMigrate(_mOrder.Order{})
	db.AutoMigrate(_mOrder.OrderDetail{})
	db.AutoMigrate(_mOrder.Payment{})
}
