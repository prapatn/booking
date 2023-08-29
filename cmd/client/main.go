package main

import (
	"booking/cmd/client/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(":1325"))
}
