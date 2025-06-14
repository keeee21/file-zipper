package controller

import (
	"net/http"
	"strconv"

	"file-zipper-api/usecase"

	"github.com/gin-gonic/gin"
)

type FileLogController struct {
	fileLogUsecase *usecase.FileLogUsecase
}

func NewFileLogController(fileLogUsecase *usecase.FileLogUsecase) *FileLogController {
	return &FileLogController{
		fileLogUsecase: fileLogUsecase,
	}
}

// ファイルのアップロード履歴を取得する
func (c *FileLogController) GetFileUploadHistory(ctx *gin.Context) {
	fileID, err := strconv.ParseUint(ctx.Param("file_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	logs, err := c.fileLogUsecase.GetFileUploadHistory(uint(fileID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, logs)
}

// ファイルのダウンロード履歴を取得する
func (c *FileLogController) GetFileDownloadHistory(ctx *gin.Context) {
	fileID, err := strconv.ParseUint(ctx.Param("file_id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file ID"})
		return
	}

	logs, err := c.fileLogUsecase.GetFileDownloadHistory(uint(fileID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, logs)
}
