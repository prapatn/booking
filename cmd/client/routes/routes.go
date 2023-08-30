package routes

import (
	"booking/cmd/client/controllers"
	"booking/pb"
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Init(e *echo.Echo) {
	conn := NewClient()
	room := controllers.RoomsClient{Client: pb.NewRoomsClient(conn)}

	e.GET("/api/v1/rooms", room.GetRooms)
	e.GET("/api/v1/rooms/bookings", room.GetRoomsWithBookings)
	e.GET("/api/v1/rooms/:id", room.GetRoom)
	e.POST("/api/v1/rooms", room.CreateRoom)
	e.PUT("/api/v1/rooms/:id", room.UpdateRoom)
	e.DELETE("/api/v1/rooms/:id", room.DeleteRoom)

	e.POST("/api/v1/users", controllers.CreateUser)
	e.POST("/api/v1/users/grade", controllers.SetUserGrade)

	e.GET("/api/v1/bookings", controllers.GetBookings)
	e.POST("/api/v1/bookings", controllers.CreateBooking)
	e.PUT("/api/v1/bookings/:id", controllers.UpdateBooking)
	e.DELETE("/api/v1/bookings/:id", controllers.DeleteBooking)
}

func NewClient() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	//defer conn.Close()
	return conn
}
