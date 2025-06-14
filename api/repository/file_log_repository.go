package repository

import (
	"time"

	"file-zipper-api/model"

	"gorm.io/gorm"
)

type FileLogRepository struct {
	db *gorm.DB
}

func NewFileLogRepository(db *gorm.DB) *FileLogRepository {
	return &FileLogRepository{db: db}
}

// アップロードログを作成
func (r *FileLogRepository) CreateUploadLog(fileID, userID uint) error {
	uploadLog := &model.UploadLog{
		FileID:    fileID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return r.db.Create(uploadLog).Error
}

// ダウンロードログを作成
func (r *FileLogRepository) CreateDownloadLog(fileID, userID uint) error {
	downloadLog := &model.DownloadLog{
		FileID:    fileID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return r.db.Create(downloadLog).Error
}

// 特定のファイルIDに対するアップロード履歴を取得
func (r *FileLogRepository) GetUploadLogsByFileID(fileID uint) ([]model.UploadLog, error) {
	var logs []model.UploadLog
	err := r.db.
		Preload("User").
		Preload("File").
		Where("file_id = ?", fileID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}

// 特定のファイルIDに対するダウンロード履歴を取得
func (r *FileLogRepository) GetDownloadLogsByFileID(fileID uint) ([]model.DownloadLog, error) {
	var logs []model.DownloadLog
	err := r.db.
		Preload("User").
		Preload("File").
		Where("file_id = ?", fileID).
		Order("created_at DESC").
		Find(&logs).Error
	return logs, err
}

// ファイルのアップロード情報を取得
func (r *FileLogRepository) GetFileUploadInfo(fileID uint) (*model.UploadLog, error) {
	var log model.UploadLog
	err := r.db.
		Preload("User").
		Preload("File").
		Where("file_id = ?", fileID).
		First(&log).Error
	return &log, err
}
