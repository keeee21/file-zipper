package model

import "time"

type RoomFile struct {
	ID     uint   `gorm:"primaryKey"`
	RoomID string `gorm:"not null;index"`
	FileID uint   `gorm:"not null;index"`

	Room DownloadRoom `gorm:"foreignKey:RoomID;references:ID;constraint:OnDelete:CASCADE"`
	File File         `gorm:"foreignKey:FileID;references:ID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
