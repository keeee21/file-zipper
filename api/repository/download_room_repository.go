package repository

import (
	"file-zipper-api/model"

	"gorm.io/gorm"
)

type IDownloadRoomRepository interface {
	GetRoomByID(roomID string) (*model.DownloadRoom, error)
	CreateRoom(room *model.DownloadRoom) error
}

type downloadRoomRepository struct {
	db *gorm.DB
}

func NewDownloadRoomRepository(db *gorm.DB) IDownloadRoomRepository {
	return &downloadRoomRepository{db: db}
}

func (r *downloadRoomRepository) GetRoomByID(roomID string) (*model.DownloadRoom, error) {
	var room model.DownloadRoom
	err := r.db.Where("id = ?", roomID).First(&room).Error
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *downloadRoomRepository) CreateRoom(room *model.DownloadRoom) error {
	return r.db.Create(room).Error
}
