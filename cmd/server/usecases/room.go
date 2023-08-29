package usecases

import (
	"booking/cmd/server/entities"
	"booking/cmd/server/entities/rooms"
	"booking/cmd/server/repository"
	"booking/pb"
	"errors"
)

func GetAllRooms(rooms *[]*pb.GetRoomResponse) error {
	err := repository.GetRooms(rooms)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func GetAllRoomsWithBookings(res *[]*pb.GetRoomWithBookingResponse) error {
	var rooms []entities.Room
	err := repository.GetRoomsWithBookings(&rooms)
	if err != nil {
		print(err)
		return err
	}

	for _, r := range rooms {
		var bookings []*pb.GetBookingResponse
		for _, booking := range r.Bookings {
			bookings = append(bookings, &pb.GetBookingResponse{
				Id:           int64(booking.ID),
				RoomId:       int64(booking.RoomsID),
				UserId:       int64(booking.UsersID),
				AmountPerson: int64(booking.AmountPerson),
				StartDate:    booking.StartDate.Format("15:04:05 01/02/2006"),
				EndDate:      booking.EndDate.Format("15:04:05 01/02/2006"),
			})
		}
		*res = append(*res, &pb.GetRoomWithBookingResponse{
			Id:            int64(r.ID),
			RoomName:      r.RoomName,
			MaximumPerson: int64(r.MaximumPerson),
			Bookings:      bookings,
		})
	}

	return nil
}

func GetRoomsByID(room *rooms.Show, id string) error {
	err := repository.GetRoomByID(room, id)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func CreateRoom(room *rooms.Form) error {
	err := repository.CreateRoom(room)
	if err != nil {
		print(err)
		return err
	}

	return nil
}

func UpdateRoom(room *rooms.Form, id string) error {
	rowAffected := repository.UpdateRoom(room, id)
	if rowAffected == 0 {
		return errors.New("Update Fail")
	}

	return nil
}

func DeleteRoom(id string) error {
	rowAffected := repository.DeleteRoom(id)
	if rowAffected == 0 {
		return errors.New("Delete Fail")
	}
	return nil
}
