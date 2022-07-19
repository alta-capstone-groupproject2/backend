package main

import (
	"lami/app/factory"
	"lami/app/middlewares"
	"lami/app/migration"
	"lami/app/routes"

	"lami/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func main() {
	var event coreapi.Client
	event.New(config.MidtransServerKey(), midtrans.Sandbox)
	// connection database
	dbConn := config.InitDB()
	// migration table
	migration.Migration(dbConn)
	// routes
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)

	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
