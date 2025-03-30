package model

import (
	"time"
)

type DownloadRoom struct {
	ID        string    `gorm:"primaryKey;type:char(26)"`
	Password  *string   `json:"password" gorm:"default:null"`
	ExpiredAt time.Time `gorm:"not null"`
	IsDeleted bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
