package repository

import (
	"file-zipper-api/model"

	"gorm.io/gorm"
)

type IFileRepository interface {
	GetFileById(id int) (model.File, error)
	CreateFile(file *model.File) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) IFileRepository {
	return &fileRepository{db: db}
}

// ファイル情報をDBから取得
func (fr *fileRepository) GetFileById(id int) (model.File, error) {
	var file model.File
	err := fr.db.First(&file, id).Error
	return file, err
}

// ファイル情報をDBに保存
func (fr *fileRepository) CreateFile(file *model.File) error {
	return fr.db.Create(file).Error
}
