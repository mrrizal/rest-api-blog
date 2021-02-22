package models

import (
	"time"
)

type UserModel struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (user UserModel) TableName() string {
	return "users"
}
