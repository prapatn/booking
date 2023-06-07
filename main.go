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

	e.Logger.Fatal(e.Start(":8080"))
}
