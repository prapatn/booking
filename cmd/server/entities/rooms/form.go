package rooms

import "github.com/go-playground/validator/v10"

type Form struct {
	RoomName      string `json:"room_name" form:"room_name" validate:"required"`
	MaximumPerson uint   `json:"maximum_person" form:"maximum_person"  validate:"required"`
}

func (Form) TableName() string {
	return "rooms"
}

func (r *Form) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
