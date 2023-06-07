package entities

type Room struct {
	ID            uint      `gorm:"primarykey" json:"id"`
	RoomName      string    `json:"room_name"`
	MaximumPerson uint      `json:"maximum_person"`
	Bookings      []Booking `gorm:"foreignKey:RoomsID" json:"bookings"`
}
