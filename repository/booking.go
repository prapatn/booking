package repository

import (
	"booking/database"
	"booking/entities"
	"booking/entities/bookings"
)

// func GetBookings(bookings *[]bookings.Show) error {
// 	return database.DB.Find(bookings).Error
// }

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
