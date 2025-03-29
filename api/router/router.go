package router

import (
	"file-zipper-api/controller"
	"file-zipper-api/gateway"
	"file-zipper-api/middleware"
	"file-zipper-api/repository"
	"file-zipper-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	googleRepo := gateway.NewGoogleAuthRepository()
	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, googleRepo)
	authHandler := controller.NewAuthHandler(authUsecase)
	e.POST("/api/auth/google", authHandler.GoogleLogin)

	fileRepo := repository.NewFileRepository(db)
	roomRepo := repository.NewDownloadRoomRepository(db)
	fileUsecase := usecase.NewFileUsecase(fileRepo, roomRepo)
	fileController := controller.NewFileController(fileUsecase)

	// 🔐 認証が必要なグループを定義
	authGroup := e.Group("/api")
	authGroup.Use(middleware.JWTMiddleware)

	// 認証付きエンドポイント
	authGroup.POST("/file-upload", fileController.UploadFile)
}
