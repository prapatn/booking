package request

import (
	"booking/cmd/server/usecases"
	"booking/pb"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
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

func (s *RoomsServer) GetRoomById(ctx context.Context, req *pb.GetRoomsByIdRequest) (*pb.GetRoomResponse, error) {
	var rooms pb.GetRoomResponse
	id := req.Id
	err := usecases.GetRoomsByID(&rooms, &id)
	if err != nil {
		print(err)
		return nil, err
	}

	return &rooms, nil
}

func (s *RoomsServer) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*emptypb.Empty, error) {
	err := usecases.CreateRoom(req)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *RoomsServer) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*emptypb.Empty, error) {
	id := req.Id
	err := usecases.UpdateRoom(req, id)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}

func (s *RoomsServer) DeleteRoom(ctx context.Context, req *pb.DeleteRoomByIdRequest) (*emptypb.Empty, error) {
	id := req.Id
	err := usecases.DeleteRoom(id)
	if err != nil {
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, nil
}
