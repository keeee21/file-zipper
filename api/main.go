package main

import (
	"file-zipper-api/config"
	"file-zipper-api/db"
	"file-zipper-api/middleware"
	"file-zipper-api/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()
	database := db.InitDB()
	defer db.CloseDB(database)

	e := echo.New()
	middleware.SetupCORS(e) // CORS 設定を適用

	router.InitRouter(e, database)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}
