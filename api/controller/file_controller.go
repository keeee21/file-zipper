package controller

import (
	"file-zipper-api/model"
	"file-zipper-api/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileController struct {
	fileUsecase usecase.IFileUsecase
}

func NewFileController(fileUsecase usecase.IFileUsecase) *FileController {
	return &FileController{fileUsecase}
}

func (fc *FileController) UploadFile(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("❌ パニックが発生:", r)
			c.JSON(http.StatusInternalServerError, map[string]string{"error": "予期しないエラーが発生しました"})
		}
	}()

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

	downloadRoomRes, err := fc.fileUsecase.CreateDownloadRoom(&fileModel, password)
	if err != nil {
		fmt.Println("❌ ダウンロードルームの作成に失敗:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "ダウンロードルームの作成に失敗しました"})
	}
	fmt.Println("✅ ダウンロードルームを作成しました:", downloadRoomRes)

	return c.JSON(http.StatusOK, uploadRes)
}
