package model

import (
	"time"
)

type DownloadRoom struct {
	ID        string    `gorm:"primaryKey;type:char(26)"`
	URL       string    `gorm:"unique;not null"`
	Password  *string   `json:"password" gorm:"default:null"`
	ExpiredAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
