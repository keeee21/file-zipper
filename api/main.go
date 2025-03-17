package main

import (
	"file-zipper-api/config"
	"file-zipper-api/db"
	"file-zipper-api/router"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// 環境変数のロード
	config.LoadEnv()

	// データベース接続
	database := db.InitDB()
	defer db.CloseDB(database)

	e := echo.New()

	// ルーターの初期化
	router.InitRouter(e, database)

	// テスト用エンドポイント
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// サーバー起動
	e.Logger.Fatal(e.Start("0.0.0.0:3001"))
}
