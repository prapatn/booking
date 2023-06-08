package repository

import (
	"booking/database"
	"booking/entities"
	"booking/entities/bookings"
)

func GetBookings(bookings *[]bookings.Show) error {
	return database.DB.Table("bookings").
		Select("bookings.amount_person, bookings.start_date, bookings.end_date,rooms.room_name,CONCAT(IFNULL(users.name_prefix,\"\") ,\" \",users.first_name,\" \",users.last_name) as full_name ").
		Joins("JOIN users on users.id = bookings.users_id").
		Joins("JOIN rooms on rooms.id = bookings.rooms_id").
		Find(&bookings).Error
}

// func GetBookingsWithRoomAndUser(bookings *[]bookings.Show) error {
// 	return database.DB.Preload("Room").Find(bookings).Error
// }

func CreateBooking(booking *bookings.Add) error {
	return database.DB.Table(booking.TableName()).Create(booking).Error
}

func UpdateBooking(booking *bookings.Update, id string) int64 {
	return database.DB.Table(booking.TableName()).Where("id = ?", id).Updates(booking).RowsAffected
}

func DeleteBooking(id string) int64 {
	return database.DB.Where("id = ?", id).Delete(&entities.Booking{}).RowsAffected
}
