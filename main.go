package main

import (
	"booking/controllers"
	"booking/database"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	database.NewDB()

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

	e.Logger.Fatal(e.Start(":8080"))
}
