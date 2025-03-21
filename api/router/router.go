package router

import (
	"file-zipper-api/controller"
	"file-zipper-api/repository"
	"file-zipper-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	fileRepo := repository.NewFileRepository(db)
	roomRepo := repository.NewDownloadRoomRepository(db)
	fileUsecase := usecase.NewFileUsecase(fileRepo, roomRepo)
	fileController := controller.NewFileController(fileUsecase)

	e.POST("/api/file-upload", fileController.UploadFile)
}
