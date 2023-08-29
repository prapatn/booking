package request

import (
	"booking/cmd/server/entities/rooms"
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
	rooms := new([]rooms.Show)

	err := usecases.GetAllRooms(rooms)
	if err != nil {
		print(err)
		return nil, err
	}

	for _, r := range *rooms {
		res = append(res, &pb.GetRoomResponse{
			Id:            int64(r.ID),
			RoomName:      r.RoomName,
			MaximumPerson: int64(r.MaximumPerson),
		})
	}
	roomsResponse := pb.GetRoomsResponse{GetRoomResponse: res}
	return &roomsResponse, nil
}
