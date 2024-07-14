package main

import (
	"PostgreSQLProject/accounts"
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ctx := context.Background()
	accountsHandler, err := accounts.New(&ctx)
	if err != nil {
		panic(err)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/patch", accountsHandler.PatchAccount)
	e.POST("/account/change", accountsHandler.ChangeAccount)
	e.POST("/account/delete", accountsHandler.DeleteAccount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
