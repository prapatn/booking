package main

import (
	"booking/cmd/client/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	routes.Init(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "AccessTime => ${time_rfc3339_nano}\n" +
			"Host => ${host}, RemoteIP => ${remote_ip},\n" +
			"Method => ${method},\n" +
			"URI => ${uri}, Status => ${status},\n" +
			"Error => ${error},\n" +
			"UserAgent => ${user_agent}\n" +
			"--------------\n",
		Output: e.Logger.Output(),
	}))

	e.Logger.Fatal(e.Start(":1325"))
}
