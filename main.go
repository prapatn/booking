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
	e.GET("/api/v1/rooms/:id", controllers.GetRoom)
	e.POST("/api/v1/rooms", controllers.CreateRoom)
	e.PUT("/api/v1/rooms/:id", controllers.UpdateRoom)
	e.DELETE("/api/v1/rooms/:id", controllers.DeleteRoom)

	e.Logger.Fatal(e.Start(":8080"))
}
