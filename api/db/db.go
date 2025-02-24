package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"file-zipper-api/config"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	dbPort, err := strconv.Atoi(config.GetEnv("DB_PORT"))
	if err != nil {
		log.Fatalf("環境変数 DB_PORT の変換エラー: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		dbPort,
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
	)

	// データベースに接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}

	// データベース接続の確認
	if err := db.Ping(); err != nil {
		log.Fatal("データベースへのPingエラー:", err)
	}

	log.Println("データベースに接続しました。")
	return db, nil
}
