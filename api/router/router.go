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
	roomFileRepo := repository.NewRoomFilesRepository(db)
	fileUsecase := usecase.NewFileUsecase(fileRepo, roomRepo, roomFileRepo)
	fileController := controller.NewFileController(fileUsecase)

	roomUsecase := usecase.NewDownloadRoomUsecase(roomRepo)
	downloadRoomController := controller.NewDownloadRoomController(roomUsecase)

	// 🔐 認証が必要なグループを定義
	authGroup := e.Group("/api")
	authGroup.Use(middleware.JWTMiddleware)

	// 認証付きエンドポイント
	authGroup.GET("/user/info", authHandler.GetUserInfo)
	authGroup.POST("/file-upload", fileController.UploadFile)
	authGroup.GET("/rooms/:roomID/validity", downloadRoomController.CheckDownloadRoomValidity)
	authGroup.GET("/files/:roomID/name", fileController.GetFileNamesByRoomId)
	authGroup.POST("/files/:roomID/signed-url", fileController.GetSignedUrl)
}
