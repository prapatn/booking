package usecases

import (
	"booking/cmd/server/entities/bookings"
	"booking/cmd/server/repository"
	"errors"
)

func GetAllBookings(bookings *[]bookings.Show) error {
	err := repository.GetBookings(bookings)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func CreateBooking(booking *bookings.Add) error {
	err := repository.CreateBooking(booking)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func UpdateBooking(booking *bookings.Update, id string) error {
	rowAffected := repository.UpdateBooking(booking, id)
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}

	return nil
}

func DeleteBooking(id string) error {
	rowAffected := repository.DeleteBooking(id)
	if rowAffected == 0 {
		return errors.New("Delete Fail")
	}
	return nil
}
