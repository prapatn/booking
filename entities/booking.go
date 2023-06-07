package entities

import (
	"time"
)

type Booking struct {
	ID           uint      `gorm:"primarykey" json:"id"`
	RoomsID      uint      `json:"rooms_id"`
	UsersID      uint      `json:"users_id"`
	AmountPerson int       `json:"amount_person"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}
