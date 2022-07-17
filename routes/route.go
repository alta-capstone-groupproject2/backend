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
	e.GET("/users/stores", presenter.UserPresenter.GetStoreSubmission, middlewares.JWTMiddleware())
	e.PUT("/users/stores/:id", presenter.UserPresenter.UpdateStatusAccount, middlewares.JWTMiddleware())

	e.POST("/cultures", presenter.CulturePresenter.PostCulture, middlewares.JWTMiddleware())
	e.GET("/cultures", presenter.CulturePresenter.GetCulture)
	e.GET("/cultures/:cultureID", presenter.CulturePresenter.GetCulturebyIDCulture)
	e.PUT("/cultures/:cultureID", presenter.CulturePresenter.PutCulture, middlewares.JWTMiddleware())
	e.DELETE("/cultures/:cultureID", presenter.CulturePresenter.DeleteCulture, middlewares.JWTMiddleware())
	
	e.POST("/cultures/reports/:cultureID", presenter.CulturePresenter.PostCultureReport, middlewares.JWTMiddleware())
	e.GET("/cultures/reports/:cultureID", presenter.CulturePresenter.GetCultureReport, middlewares.JWTMiddleware())

	e.POST("/events/comments", presenter.CommentPresenter.Add, middlewares.JWTMiddleware())
	e.GET("/events/comments/:id", presenter.CommentPresenter.Get, middlewares.JWTMiddleware())

	e.GET("/events", presenter.EventPresenter.GetAll)
	e.GET("/events/:id", presenter.EventPresenter.GetDataById)
	e.POST("/events", presenter.EventPresenter.InsertData, middlewares.JWTMiddleware())
	e.PUT("/events/:id", presenter.EventPresenter.UpdateData, middlewares.JWTMiddleware())
	e.DELETE("/events/:id", presenter.EventPresenter.DeleteData, middlewares.JWTMiddleware())
	e.GET("/myevents", presenter.EventPresenter.GetEventByUser, middlewares.JWTMiddleware())

	e.POST("/events/participations", presenter.ParticipantPresenter.Joined, middlewares.JWTMiddleware())
	e.GET("/events/participations", presenter.ParticipantPresenter.GetAllEventParticipant, middlewares.JWTMiddleware())
	e.DELETE("/events/participations/:id", presenter.ParticipantPresenter.DeleteEventbyParticipant, middlewares.JWTMiddleware())

	// Product
	e.POST("/products", presenter.ProductPresenter.PostProduct, middlewares.JWTMiddleware())
	e.PUT("/products/:productID", presenter.ProductPresenter.PutProduct, middlewares.JWTMiddleware())
	e.DELETE("/products/:productID", presenter.ProductPresenter.DeleteProduct, middlewares.JWTMiddleware())
	e.GET("/products/:productID", presenter.ProductPresenter.GetProductbyIDProduct)
	e.GET("/users/products", presenter.ProductPresenter.GetMyProduct, middlewares.JWTMiddleware())

	// Rating
	e.POST("/products/ratings/:productID", presenter.ProductPresenter.PostProductRating, middlewares.JWTMiddleware())
	e.GET("/products/ratings/:productID", presenter.ProductPresenter.GetProductRating)

	// Cart
	e.POST("/carts", presenter.CartPresenter.PostCart, middlewares.JWTMiddleware())
	e.GET("/carts", presenter.CartPresenter.GetCart, middlewares.JWTMiddleware())
	e.PUT("/carts/:cartID", presenter.CartPresenter.PutCart, middlewares.JWTMiddleware())
	e.DELETE("carts/:cartID", presenter.CartPresenter.DeletedCart, middlewares.JWTMiddleware())

	return e
}
