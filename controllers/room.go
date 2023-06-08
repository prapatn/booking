package controllers

import (
	"booking/entities/rooms"
	"booking/usecases"
	"net/http"

	"github.com/labstack/echo"
)

func GetRooms(c echo.Context) error {
	rooms := new([]rooms.Show)
	err := usecases.GetAllRooms(rooms)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

func GetRoom(c echo.Context) error {
	room := new(rooms.Show)
	id := c.Param("id")
	err := usecases.GetRoomsByID(room, id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}

func CreateRoom(c echo.Context) error {
	room := new(rooms.Form)
	err := c.Bind(room)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = room.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = usecases.CreateRoom(room)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func UpdateRoom(c echo.Context) error {
	room := new(rooms.Form)
	err := c.Bind(room)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = room.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	id := c.Param("id")
	err = usecases.UpdateRoom(room, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func DeleteRoom(c echo.Context) error {
	id := c.Param("id")
	err := usecases.DeleteRoom(id)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
