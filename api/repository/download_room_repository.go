package repository

import (
	"file-zipper-api/model"

	"gorm.io/gorm"
)

type IDownloadRoomRepository interface {
	CreateRoom(room *model.DownloadRoom) error
}

type downloadRoomRepository struct {
	db *gorm.DB
}

func NewDownloadRoomRepository(db *gorm.DB) IDownloadRoomRepository {
	return &downloadRoomRepository{db: db}
}

func (r *downloadRoomRepository) CreateRoom(room *model.DownloadRoom) error {
	return r.db.Create(room).Error
}
