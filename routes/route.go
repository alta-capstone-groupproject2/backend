package routes

import (
	"lami/app/factory"
	"lami/app/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Pre(middlewares.RemoveTrailingSlash())

	e.Use(middlewares.CorsMiddleware())

	e.POST("/register", presenter.UserPresenter.Insert)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/users", presenter.UserPresenter.GetDataById, middlewares.JWTMiddleware())
	e.PUT("/users", presenter.UserPresenter.Update, middlewares.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.Delete, middlewares.JWTMiddleware())
	e.POST("/users/stores", presenter.UserPresenter.AccountUpgrade, middlewares.JWTMiddleware())

	e.GET("/stores/submissions", presenter.UserPresenter.GetStoreSubmission, middlewares.JWTMiddleware())
	e.PUT("/stores/submissions/:id", presenter.UserPresenter.UpdateStatusAccount, middlewares.JWTMiddleware())

	e.POST("/users/verify", presenter.UserPresenter.GmailVerification)
	e.GET("/users/confirm/:encrypt", presenter.UserPresenter.InsertFromVerificaton)

	e.POST("/cultures", presenter.CulturePresenter.PostCulture, middlewares.JWTMiddleware())
	e.GET("/cultures", presenter.CulturePresenter.GetCulture)
	e.GET("/cultures/:cultureID", presenter.CulturePresenter.GetCulturebyIDCulture)
	e.PUT("/cultures/:cultureID", presenter.CulturePresenter.PutCulture, middlewares.JWTMiddleware())
	e.DELETE("/cultures/:cultureID", presenter.CulturePresenter.DeleteCulture, middlewares.JWTMiddleware())

	e.POST("/cultures/reports/:cultureID", presenter.CulturePresenter.PostCultureReport, middlewares.JWTMiddleware())
	e.GET("/cultures/reports/:cultureID", presenter.CulturePresenter.GetCultureReport, middlewares.JWTMiddleware())

	//event data
	e.POST("/events", presenter.EventPresenter.InsertData, middlewares.JWTMiddleware())
	e.GET("/events", presenter.EventPresenter.GetAll)
	e.GET("/events/:id", presenter.EventPresenter.GetDataById)
	e.DELETE("/events/:id", presenter.EventPresenter.DeleteData, middlewares.JWTMiddleware())
	e.GET("/users/events", presenter.EventPresenter.GetEventByUser, middlewares.JWTMiddleware())

	//Attendee Event
	e.GET("/events/attendees/:id", presenter.EventPresenter.GetEventAttendeesData, middlewares.JWTMiddleware())

	e.POST("/events/comments", presenter.CommentPresenter.Add, middlewares.JWTMiddleware())
	e.GET("/events/comments/:id", presenter.CommentPresenter.Get)

	e.POST("/events/participations", presenter.ParticipantPresenter.Joined, middlewares.JWTMiddleware())
	e.GET("/events/participations", presenter.ParticipantPresenter.GetAllEventParticipant, middlewares.JWTMiddleware())
	e.DELETE("/events/participations/:id", presenter.ParticipantPresenter.DeleteEventbyParticipant, middlewares.JWTMiddleware())

	//Payment Event
	e.POST("/events/payments", presenter.ParticipantPresenter.CreatePayment, middlewares.JWTMiddleware())
	e.GET("/events/payment_details", presenter.ParticipantPresenter.GetDetailPayment, middlewares.JWTMiddleware())
	e.GET("/events/payments/status", presenter.ParticipantPresenter.CheckStatusPayment, middlewares.JWTMiddleware())
	//Midtrans Web Hook
	e.POST("/events/payments/webhook", presenter.ParticipantPresenter.MidtransWebHook)

	//submission by user
	e.GET("/events/submissions", presenter.EventPresenter.GetSubmissionAll, middlewares.JWTMiddleware())
	e.GET("/events/submissions/:id", presenter.EventPresenter.GetSubmissionByID, middlewares.JWTMiddleware())
	e.PUT("/events/submissions/:id", presenter.EventPresenter.UpdateData, middlewares.JWTMiddleware())

	// Product
	e.POST("/products", presenter.ProductPresenter.PostProduct, middlewares.JWTMiddleware())
	e.GET("/products", presenter.ProductPresenter.GetProductList)
	e.PUT("/products/:productID", presenter.ProductPresenter.PutProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:productID", presenter.ProductPresenter.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/products/:productID", presenter.ProductPresenter.GetProductbyIDProduct)
	e.GET("/users/products", presenter.ProductPresenter.GetMyProduct, middlewares.JWTMiddleware())

	// Rating
	e.POST("/products/ratings/:productID", presenter.ProductPresenter.PostProductRating, middlewares.JWTMiddleware())
	e.GET("/products/ratings/:productID", presenter.ProductPresenter.GetProductRating)

	// Order
	e.POST("/orders", presenter.OrderPresenter.PostOrder, middlewares.JWTMiddleware())

	// Cart
	e.POST("/carts", presenter.CartPresenter.PostCart, middlewares.JWTMiddleware())
	e.GET("/carts", presenter.CartPresenter.GetCart, middlewares.JWTMiddleware())
	e.PUT("/carts/:cartID", presenter.CartPresenter.PutCart, middlewares.JWTMiddleware())
	e.DELETE("carts/:cartID", presenter.CartPresenter.DeletedCart, middlewares.JWTMiddleware())

	// PaymentOrder
	e.POST("/payments/:type", presenter.PaymentPresenter.PostPayment, middlewares.JWTMiddleware())
	e.PUT("/payments/confirm", presenter.PaymentPresenter.PutPayment, middlewares.JWTMiddleware())

	return e
}
