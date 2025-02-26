package main

import (
	"file-zipper-api/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// 環境変数のロード
	config.LoadEnv()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}
