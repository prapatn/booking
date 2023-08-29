package request

import (
	"booking/cmd/server/usecases"
	"booking/pb"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type RoomsServer struct {
	pb.UnsafeRoomsServer
}

func (s *RoomsServer) GetRooms(ctx context.Context, emp *emptypb.Empty) (*pb.GetRoomsResponse, error) {
	var res []*pb.GetRoomResponse
	err := usecases.GetAllRooms(&res)
	if err != nil {
		print(err)
		return nil, err
	}

	roomsResponse := pb.GetRoomsResponse{GetRoomResponse: res}
	return &roomsResponse, nil
}

func (s *RoomsServer) GetRoomsWithBookings(ctx context.Context, emp *emptypb.Empty) (*pb.GetRoomsWithBookingsResponse, error) {
	var rooms []*pb.GetRoomWithBookingResponse
	err := usecases.GetAllRoomsWithBookings(&rooms)
	if err != nil {
		print(err)
		return nil, err
	}

	res := pb.GetRoomsWithBookingsResponse{Rooms: rooms}
	return &res, nil
}
