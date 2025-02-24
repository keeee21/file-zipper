package main

import (
	"file-zipper-api/config"
	"file-zipper-api/db"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// 環境変数のロード
	config.LoadEnv()

	// DB接続
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to DB: %v", err)
	}
	defer database.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}
