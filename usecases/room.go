package usecases

import (
	"booking/entities"
	"booking/entities/rooms"
	"booking/repository"
	"errors"
)

func GetAllRooms(rooms *[]rooms.Show) error {
	err := repository.GetRooms(rooms)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetAllRoomsWithBookings(rooms *[]entities.Room) error {
	err := repository.GetRoomsWithBookings(rooms)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetRoomsByID(room *rooms.Show, id string) error {
	err := repository.GetRoomByID(room, id)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func CreateRoom(room *rooms.Form) error {
	err := repository.CreateRoom(room)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func UpdateRoom(room *rooms.Form, id string) error {
	rowAffected := repository.UpdateRoom(room, id)
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}

	return nil
}

func DeleteRoom(id string) error {
	rowAffected := repository.DeleteRoom(id)
	if rowAffected == 0 {
		return errors.New("Delete Fail")
	}
	return nil
}
