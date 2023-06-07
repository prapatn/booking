package controllers

import (
	"booking/entities"
	"booking/usecases"
	"net/http"

	"github.com/labstack/echo"
)

func GetRooms(c echo.Context) error {
	rooms := new([]entities.Room)
	err := usecases.GetAllRooms(rooms)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}
