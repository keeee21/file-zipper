package repository

import (
	"gorm.io/gorm"
)

type IRoomFilesRepository interface {
	GetFileIdsByRoomId(roomID string) ([]int, error)
	Create(roomID string, fileID uint) error
}

type roomFilesRepository struct {
	db *gorm.DB
}

func NewRoomFilesRepository(db *gorm.DB) IRoomFilesRepository {
	return &roomFilesRepository{db: db}
}

func (r *roomFilesRepository) Create(roomID string, fileID uint) error {
	err := r.db.Table("room_files").Create(map[string]interface{}{
		"room_id":    roomID,
		"file_id":    fileID,
		"created_at": gorm.Expr("NOW()"),
		"updated_at": gorm.Expr("NOW()"),
	}).Error
	if err != nil {
		return err
	}
	return nil
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
