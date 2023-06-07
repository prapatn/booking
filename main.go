package main

import (
	database "booking/db"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	database.NewDB()

	e.Logger.Fatal(e.Start(":8080"))
}
