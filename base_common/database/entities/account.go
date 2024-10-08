package entities

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	UserName  string         `json:"user_name"`
	Password  string         `json:"password"`
	ID        int64          `json:"id"`
	Status    int            `json:"status"`
	Role      int            `json:"role"`
}
