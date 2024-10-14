package entities

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Sex         string         `json:"sex"`
	RoomNumber  string         `json:"room_number"`
	DoctorName  string         `json:"doctor_name"`
	Status      string         `json:"status"`
	OrderNumber int            `json:"order_number"`
	ID          int64          `json:"id"`
	Age         int            `json:"age"`
}
