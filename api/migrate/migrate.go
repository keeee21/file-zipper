package main

import (
	"file-zipper-api/db"
	"file-zipper-api/model"
	"fmt"
)

func main() {
	dbConn := db.InitDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.File{}, &model.DownloadRoom{}, &model.RoomFile{})
}
