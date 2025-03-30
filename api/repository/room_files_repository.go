package repository

import (
	"gorm.io/gorm"
)

type IRoomFilesRepository interface {
	GetFileIdsByRoomId(roomID string) ([]int, error)
}

type roomFilesRepository struct {
	db *gorm.DB
}

func NewRoomFilesRepository(db *gorm.DB) IRoomFilesRepository {
	return &roomFilesRepository{db: db}
}

func (r *roomFilesRepository) GetFileIdsByRoomId(roomID string) ([]int, error) {
	var fileIDs []int
	err := r.db.Table("room_files").
		Select("file_id").
		Where("room_id = ?", roomID).
		Pluck("file_id", &fileIDs).Error
	if err != nil {
		return nil, err
	}
	return fileIDs, nil
}
