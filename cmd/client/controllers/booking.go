package controllers

import (
	"booking/cmd/server/entities/bookings"
	"booking/cmd/server/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBookings(c echo.Context) error {
	bookings := new([]bookings.Show)
	err := usecases.GetAllBookings(bookings)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, bookings)
}

func CreateBooking(c echo.Context) error {
	booking := new(bookings.Add)
	err := c.Bind(booking)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = booking.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = usecases.CreateBooking(booking)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func UpdateBooking(c echo.Context) error {
	booking := new(bookings.Update)
	err := c.Bind(booking)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = booking.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	err = usecases.UpdateBooking(booking, id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusCreated, "Success")
}

func DeleteBooking(c echo.Context) error {
	id := c.Param("id")
	err := usecases.DeleteBooking(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
