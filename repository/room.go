package repository

import (
	"booking/database"
	"booking/entities"
)

func GetRooms(rooms *[]entities.Room) error {
	return database.DB.Preload("Bookings").Find(rooms).Error
}
