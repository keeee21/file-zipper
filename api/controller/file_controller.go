package controller

import (
	"file-zipper-api/model"
	"file-zipper-api/usecase"
	"file-zipper-api/util"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FileController struct {
	fileUsecase    usecase.IFileUsecase
	fileLogUsecase *usecase.FileLogUsecase
}

func NewFileController(fileUsecase usecase.IFileUsecase, fileLogUsecase *usecase.FileLogUsecase) *FileController {
	return &FileController{
		fileUsecase:    fileUsecase,
		fileLogUsecase: fileLogUsecase,
	}
}

func (fc *FileController) UploadFile(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❌ パニックが発生:", r)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "予期しないエラーが発生しました"})
		}
	}()

	// ユーザーIDを取得
	userID, err := util.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "認証が必要です"})
	}

	expirationStr := c.FormValue("expiration")
	expirationDays, err := strconv.Atoi(expirationStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "有効期限の取得に失敗しました"})
	}

	// フォームデータ取得
	fileHeader, err := c.FormFile("file")
	if err != nil {
		fmt.Println("❌ ファイルが選択されていません")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ファイルが必要です"})
	}

	// ファイルオープン
	file, err := fileHeader.Open()
	if err != nil {
		fmt.Println("❌ ファイルを開けませんでした:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ファイルを開けませんでした"})
	}
	defer file.Close()

	// ファイルのバイナリデータ取得
	fileData := make([]byte, fileHeader.Size)
	_, err = file.Read(fileData)
	if err != nil {
		fmt.Println("❌ ファイルの読み込みに失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ファイルを読み込めませんでした"})
	}

	// フォームからパスワードを取得
	password := c.FormValue("password")

	// ファイル情報作成
	fileModel := model.File{}

	// ✅ ユースケースを呼び出し（エラーハンドリング追加）
	uploadRes, err := fc.fileUsecase.Upload(&fileModel, fileData, fileHeader.Filename)
	if err != nil {
		fmt.Println("❌ ファイルのアップロードに失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ファイルのアップロードに失敗しました"})
	}

	// アップロードログを記録
	err = fc.fileLogUsecase.LogFileUpload(uploadRes.ID, userID)
	if err != nil {
		fmt.Println("❌ アップロードログの記録に失敗:", err)
		// ログの記録失敗は致命的ではないので、エラーを返さない
	}

	downloadRoomRes, err := fc.fileUsecase.CreateDownloadRoom(&fileModel, password, expirationDays)
	if err != nil {
		fmt.Println("❌ ダウンロードルームの作成に失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ダウンロードルームの作成に失敗しました"})
	}

	// ダウンロードルームIDと、fileIDを使って、room_filesテーブルにレコードを作成
	err = fc.fileUsecase.CreateRoomFile(downloadRoomRes.ID, uploadRes.ID)
	if err != nil {
		fmt.Println("❌ room_filesテーブルの作成に失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "room_filesテーブルの作成に失敗しました"})
	}

	res := map[string]interface{}{
		"name": uploadRes.Name,
		"url":  os.Getenv("FRONTEND_ORIGIN") + downloadRoomRes.ID,
	}

	return c.JSON(http.StatusOK, res)
}

func (fc *FileController) GetFileNamesByRoomId(c echo.Context) error {
	roomId := c.Param("roomID")

	fileNames, err := fc.fileUsecase.GetFileNamesByRoomId(roomId)
	if err != nil {
		fmt.Println("❌ ファイル名の取得に失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ファイル名の取得に失敗しました"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": map[string][]string{"fileNames": fileNames},
	})
}

type SignedUrlRequest struct {
	Password string `json:"password"`
}

func (fc *FileController) GetSignedUrl(c echo.Context) error {
	// ユーザーIDを取得
	userID, err := util.GetUserID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "認証が必要です"})
	}

	roomId := c.Param("roomID")
	signedUrls := make([]string, 0)

	files, err := fc.fileUsecase.GetFileByRoomId(roomId)
	if err != nil {
		fmt.Println("❌ ファイルの取得に失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ファイルの取得に失敗しました",
		})
	}
	if len(files) == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "roomIdに関連付けられたファイルがありません",
		})
	}

	var req SignedUrlRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "リクエストが不正です",
		})
	}

	ok, err := fc.fileUsecase.VerifyRoomPassword(roomId, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "パスワード検証に失敗しました",
		})
	}

	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "パスワードが間違っています",
		})
	}

	for _, file := range files {
		url, err := fc.fileUsecase.GetSignedUrl(file.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "サイン付きURLの取得に失敗しました",
			})
		}
		signedUrls = append(signedUrls, url)

		// ダウンロードログを記録
		err = fc.fileLogUsecase.LogFileDownload(file.ID, userID)
		if err != nil {
			fmt.Println("❌ ダウンロードログの記録に失敗:", err)
			// ログの記録失敗は致命的ではないので、エラーを返さない
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": signedUrls,
	})
}
