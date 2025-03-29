package model

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key"`
	GoogleSub string    `json:"google_sub" gorm:"unique"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	GoogleSub string `json:"google_sub"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Token     string `json:"token"`
}
