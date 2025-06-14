package usecase

import (
	"file-zipper-api/model"
	"file-zipper-api/repository"
)

type FileLogUsecase struct {
	fileLogRepo *repository.FileLogRepository
}

func NewFileLogUsecase(fileLogRepo *repository.FileLogRepository) *FileLogUsecase {
	return &FileLogUsecase{
		fileLogRepo: fileLogRepo,
	}
}

// ファイルアップロードのログを記録
func (u *FileLogUsecase) LogFileUpload(fileID, userID uint) error {
	return u.fileLogRepo.CreateUploadLog(fileID, userID)
}

// ファイルダウンロードのログを記録
func (u *FileLogUsecase) LogFileDownload(fileID, userID uint) error {
	return u.fileLogRepo.CreateDownloadLog(fileID, userID)
}

// ファイルのアップロード履歴を取得
func (u *FileLogUsecase) GetFileUploadHistory(fileID uint) ([]model.UploadLog, error) {
	return u.fileLogRepo.GetUploadLogsByFileID(fileID)
}

// ファイルのダウンロード履歴を取得
func (u *FileLogUsecase) GetFileDownloadHistory(fileID uint) ([]model.DownloadLog, error) {
	return u.fileLogRepo.GetDownloadLogsByFileID(fileID)
}

// ファイルのアップロード情報を取得
func (u *FileLogUsecase) GetFileUploadInfo(fileID uint) (*model.UploadLog, error) {
	return u.fileLogRepo.GetFileUploadInfo(fileID)
}
