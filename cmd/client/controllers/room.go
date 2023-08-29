package controllers

import (
	"booking/cmd/server/entities/rooms"
	"booking/cmd/server/usecases"
	"booking/pb"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/labstack/echo/v4"
)

type RoomsClient struct {
	Client pb.RoomsClient
}

func (r RoomsClient) GetRooms(c echo.Context) error {
	rooms, err := r.Client.GetRooms(c.Request().Context(), &empty.Empty{})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, rooms.GetRoomResponse)
}

func (r RoomsClient) GetRoomsWithBookings(c echo.Context) error {
	rooms, err := r.Client.GetRoomsWithBookings(c.Request().Context(), &empty.Empty{})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, rooms.Rooms)
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
		return c.String(http.StatusBadRequest, err.Error())
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
