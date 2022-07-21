package factory

import (
	_userBusiness "lami/app/features/users/business"
	_userData "lami/app/features/users/data"
	_userPresentation "lami/app/features/users/presentation"

	_authBusiness "lami/app/features/auth/business"
	_authData "lami/app/features/auth/data"
	_authPresentation "lami/app/features/auth/presentation"

	_eventBusiness "lami/app/features/events/business"
	_eventData "lami/app/features/events/data"
	_eventPresentation "lami/app/features/events/presentation"

	_commentBusiness "lami/app/features/comments/business"
	_commentData "lami/app/features/comments/data"
	_commentPresentation "lami/app/features/comments/presentation"

	_participantBusiness "lami/app/features/participants/business"
	_participantData "lami/app/features/participants/data"
	_participantPresentation "lami/app/features/participants/presentation"

	_productBusiness "lami/app/features/products/business"
	_productData "lami/app/features/products/data"
	_productPresentation "lami/app/features/products/presentation"

	_cultureBusiness "lami/app/features/cultures/business"
	_cultureData "lami/app/features/cultures/data"
	_culturePresentation "lami/app/features/cultures/presentation"

	_cartBusiness "lami/app/features/carts/business"
	_cartData "lami/app/features/carts/data"
	_cartPresentation "lami/app/features/carts/presentation"

	_orderBusiness "lami/app/features/orders/business"
	_orderData "lami/app/features/orders/data"
	_orderPresentation "lami/app/features/orders/presentation"

	_paymentBusiness "lami/app/features/paymentsorder/business"
	_paymentData "lami/app/features/paymentsorder/data"
	_paymentPresentation "lami/app/features/paymentsorder/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter        *_userPresentation.UserHandler
	AuthPresenter        *_authPresentation.AuthHandler
	EventPresenter       *_eventPresentation.EventHandler
	ParticipantPresenter *_participantPresentation.ParticipantHandler
	CommentPresenter     *_commentPresentation.CommentHandler
	ProductPresenter     *_productPresentation.ProductHandler
	CartPresenter        *_cartPresentation.CartHandler
	CulturePresenter     *_culturePresentation.CultureHandler
	OrderPresenter       *_orderPresentation.OrderHandler
	PaymentPresenter     *_paymentPresentation.PaymentsHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {

	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	authData := _authData.NewAuthRepository(dbConn)
	authBusiness := _authBusiness.NewAuthBusiness(authData)
	authPresentation := _authPresentation.NewAuthHandler(authBusiness)

	eventData := _eventData.NewEventRepository(dbConn)
	eventBusiness := _eventBusiness.NewEventBusiness(eventData)
	eventPresentation := _eventPresentation.NewEventHandler(eventBusiness)

	commentData := _commentData.NewCommentRepository(dbConn)
	commentBusiness := _commentBusiness.NewCommentBusiness(commentData)
	commentPresentation := _commentPresentation.NewCommentHandler(commentBusiness)

	participantData := _participantData.NewParticipantRepository(dbConn)
	participantBusiness := _participantBusiness.NewParticipantBusiness(participantData)
	participantPresentation := _participantPresentation.NewParticipantHandler(participantBusiness)

	cultureData := _cultureData.NewCultureRepository(dbConn)
	cultureBusiness := _cultureBusiness.NewCultureBusiness(cultureData)
	culturePresentation := _culturePresentation.NewCultureHandler(cultureBusiness)

	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)

	cartData := _cartData.NewCartRepository(dbConn)
	cartBusiness := _cartBusiness.NewCartBusiness(cartData)
	cartPresentation := _cartPresentation.NewCartHandler(cartBusiness)

	orderData := _orderData.NewOrderRepository(dbConn)
	orderBusiness := _orderBusiness.NewOrderBusiness(orderData)
	orderPresentation := _orderPresentation.NewOrderHandler(orderBusiness)

	paymentData := _paymentData.NewPaymentsRepository(dbConn)
	paymentBusiness := _paymentBusiness.NewPaymentsBusiness(paymentData)
	paymentPresentation := _paymentPresentation.NewPaymentHandler(paymentBusiness)

	return Presenter{
		UserPresenter:        userPresentation,
		AuthPresenter:        authPresentation,
		EventPresenter:       eventPresentation,
		ParticipantPresenter: participantPresentation,
		CommentPresenter:     commentPresentation,
		ProductPresenter:     productPresentation,
		CartPresenter:        cartPresentation,
		CulturePresenter:     culturePresentation,
		OrderPresenter:       orderPresentation,
		PaymentPresenter:     paymentPresentation,
	}
}
