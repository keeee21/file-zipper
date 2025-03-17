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
	fileUsecase := usecase.NewFileUsecase(fileRepo)
	fileController := controller.NewFileController(fileUsecase)

	e.POST("/api/file-upload", fileController.UploadFile)
}
