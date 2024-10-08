package entities

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        int64          `json:"id"`
	UserName  string         `json:"user_name"`
	Password  string         `json:"password"`
	Status    int            `json:"status"`
	Role      int            `json:"role"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
