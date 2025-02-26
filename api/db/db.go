package db

import (
	"file-zipper-api/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config.LoadEnv()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s", config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"), config.GetEnv("DB_HOST"), config.GetEnv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("データベース接続エラー:", err)
	}

	log.Println("データベースに接続しました。")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB closed")
}
