package bookings

import (
	"time"
)

type Show struct {
	RoomName     string    `json:"room_name"`
	FullName     string    `json:"full_name"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	AmountPerson int       `json:"amount_person"`
}

func (Show) TableName() string {
	return "bookings"
}
