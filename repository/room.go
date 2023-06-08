package repository

import (
	"booking/database"
	"booking/entities"
	"booking/entities/rooms"
)

func GetRooms(rooms *[]rooms.Show) error {
	return database.DB.Model(&entities.Room{}).Find(rooms).Error
}

func GetRoomsWithBookings(rooms *[]entities.Room) error {
	return database.DB.Preload("Bookings").Find(rooms).Error
}

func GetRoomByID(room *rooms.Show, id string) error {
	return database.DB.Model(&entities.Room{}).First(room, id).Error
}

func CreateRoom(room *rooms.Form) error {
	return database.DB.Table(room.TableName()).Create(room).Error
}

func UpdateRoom(room *rooms.Form, id string) int64 {
	return database.DB.Table(room.TableName()).Where("id = ?", id).Updates(room).RowsAffected
}

func DeleteRoom(id string) int64 {
	return database.DB.Where("id = ?", id).Delete(&entities.Room{}).RowsAffected
}
