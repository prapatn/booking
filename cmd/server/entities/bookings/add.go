package bookings

import "github.com/go-playground/validator/v10"

type Add struct {
	RoomsID      uint   `json:"rooms_id" form:"rooms_id" validate:"required"`
	UsersID      uint   `json:"users_id" form:"users_id" validate:"required"`
	AmountPerson int    `json:"amount_person" form:"amount_person" validate:"required"`
	StartDate    string `json:"start_date" form:"start_date" validate:"required"`
	EndDate      string `json:"end_date" form:"end_date" validate:"required"`
}

func (Add) TableName() string {
	return "bookings"
}

func (a *Add) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
