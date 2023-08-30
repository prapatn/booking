package repository

import (
	"booking/cmd/server/database"
	"booking/cmd/server/entities"
	"booking/pb"
)

func GetRooms(rooms *[]*pb.GetRoomResponse) error {
	return database.DB.Model(&entities.Room{}).Find(rooms).Error
}

func GetRoomsWithBookings(rooms *[]entities.Room) error {
	return database.DB.Preload("Bookings").Find(rooms).Error
}

func GetRoomByID(room *pb.GetRoomResponse, id *string) error {
	return database.DB.Model(&entities.Room{}).First(room, id).Error
}

func CreateRoom(room *pb.CreateRoomRequest) error {
	return database.DB.Table("rooms").Create(room).Error
}

func UpdateRoom(room *pb.UpdateRoomRequest, id string) int64 {
	return database.DB.Table("rooms").Where("id = ?", id).Updates(room).RowsAffected
}

func DeleteRoom(id string) int64 {
	return database.DB.Where("id = ?", id).Delete(&entities.Room{}).RowsAffected
}
