package usecases

import (
	"booking/cmd/server/entities/users"
	"booking/cmd/server/repository"

	"errors"
)

func CreateUser(user *users.Register) error {
	err := repository.CreateUser(user)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func UpdateGrade(user *users.Grade) error {
	id := user.ID
	rowAffected := repository.UpdateUser(user, id)
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}

	return nil
}
