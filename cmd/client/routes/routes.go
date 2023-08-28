package routes

import (
	"booking/cmd/client/controllers"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {

	e.GET("/api/v1/rooms", controllers.GetRooms)
	e.GET("/api/v1/rooms/bookings", controllers.GetRoomsWithBookings)
	e.GET("/api/v1/rooms/:id", controllers.GetRoom)
	e.POST("/api/v1/rooms", controllers.CreateRoom)
	e.PUT("/api/v1/rooms/:id", controllers.UpdateRoom)
	e.DELETE("/api/v1/rooms/:id", controllers.DeleteRoom)

	e.POST("/api/v1/users", controllers.CreateUser)
	e.POST("/api/v1/users/grade", controllers.SetUserGrade)

	e.GET("/api/v1/bookings", controllers.GetBookings)
	e.POST("/api/v1/bookings", controllers.CreateBooking)
	e.PUT("/api/v1/bookings/:id", controllers.UpdateBooking)
	e.DELETE("/api/v1/bookings/:id", controllers.DeleteBooking)
}
