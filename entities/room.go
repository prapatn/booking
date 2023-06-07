package entities

type Room struct {
	ID            uint   `gorm:"primarykey"`
	RoomName      string `json:"room_name"`
	MaximumPerson uint   `json:"maximum_person"`
	Booking       []Booking
}
