package model

import "time"

type File struct {
	ID          int       `json:"id" gorm:"primary_key"`
	FilePath    string    `json:"file_path"`
	Password    string    `json:"password"`
	ExpiredAt   time.Time `json:"expired_at"`
	DownloadUrl string    `json:"download_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FileResponse struct {
	ID          int    `json:"id" gorm:"primary_key"`
	DownloadUrl string `json:"download_url"`
}
