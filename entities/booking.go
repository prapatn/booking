package entities

import (
	"time"
)

type Booking struct {
	ID           uint      `gorm:"primarykey"`
	RoomID       uint      `json:"users_id"`
	UserID       uint      `json:"rooms_id"`
	AmountPerson uint      `json:"amount_person"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}
