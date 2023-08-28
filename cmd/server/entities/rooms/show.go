package rooms

type Show struct {
	ID            uint   `gorm:"primarykey" json:"id"`
	RoomName      string `json:"room_name"`
	MaximumPerson uint   `json:"maximum_person"`
}
