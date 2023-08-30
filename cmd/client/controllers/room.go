package controllers

import (
	"booking/cmd/server/entities/rooms"
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
func (r RoomsClient) GetRoom(c echo.Context) error {
	id := c.Param("id")
	room, err := r.Client.GetRoomById(c.Request().Context(), &pb.GetRoomsByIdRequest{Id: id})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, room)
}

func (r RoomsClient) CreateRoom(c echo.Context) error {
	room := new(rooms.Form)
	err := c.Bind(room)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = room.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err = r.Client.CreateRoom(
		c.Request().Context(),
		&pb.CreateRoomRequest{
			RoomName:      room.RoomName,
			MaximumPerson: int64(room.MaximumPerson),
		},
	)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func (r RoomsClient) UpdateRoom(c echo.Context) error {
	room := new(rooms.Form)
	err := c.Bind(room)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = room.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err = r.Client.UpdateRoom(
		c.Request().Context(),
		&pb.UpdateRoomRequest{
			Id:            c.Param("id"),
			RoomName:      room.RoomName,
			MaximumPerson: int64(room.MaximumPerson),
		},
	)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusCreated, "Success")
}

func (r RoomsClient) DeleteRoom(c echo.Context) error {
	_, err := r.Client.DeleteRoom(
		c.Request().Context(),
		&pb.DeleteRoomByIdRequest{
			Id: c.Param("id"),
		},
	)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, "Success")
}
