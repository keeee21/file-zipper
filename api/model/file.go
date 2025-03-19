package model

import "time"

type File struct {
	ID           int       `json:"id" gorm:"primary_key"`
	Name         string    `json:"name"`
	OriginalName string    `json:"original_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FileResponse struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
