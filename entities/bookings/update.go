package bookings

import (
	"github.com/go-playground/validator/v10"
)

type Update struct {
	AmountPerson int    `json:"amount_person" form:"amount_person" validate:"required"`
	StartDate    string `json:"start_date" form:"start_date" validate:"required"`
	EndDate      string `json:"end_date" form:"end_date" validate:"required"`
}

func (Update) TableName() string {
	return "bookings"
}

func (a *Update) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}
