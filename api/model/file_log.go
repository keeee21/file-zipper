package model

import (
	"time"
)

type UploadLog struct {
	FileID    uint      `gorm:"not null" json:"file_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	File      File      `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

type DownloadLog struct {
	FileID    uint      `gorm:"not null" json:"file_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	User      User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	File      File      `gorm:"foreignKey:FileID" json:"file,omitempty"`
}
