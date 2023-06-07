package usecases

import (
	"booking/entities"
	"booking/repository"
)

func GetAllRooms(rooms *[]entities.Room) error {
	err := repository.GetRooms(rooms)
	if err != nil {
		print(err)
		return err
	}

	return nil
}
