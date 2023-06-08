package users

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Register struct {
	NamePrefix   string `json:"name_prefix" form:"name_prefix" validate:"required"`
	FirstName    string `json:"first_name" form:"first_name" validate:"required"`
	LastName     string `json:"last_name" form:"last_name" validate:"required"`
	Username     string `json:"username" form:"username" validate:"required"`
	Password     string `json:"password" form:"password" validate:"required" gorm:"-"`
	PasswordHash string `json:"password_hash" form:"password_hash" `
}

func (Register) TableName() string {
	return "users"
}

func (r *Register) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *Register) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(r.Password), 14)
	if err != nil {
		return err
	}
	r.PasswordHash = string(bytes)
	return nil
}

func (r *Register) CheckPasswordHash(hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(r.Password))
	return err
}
