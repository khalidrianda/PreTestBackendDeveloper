package main

import (
	"testcase/features/main/delivery"
	"testcase/features/main/entity"
	"testcase/features/main/services"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var srv entity.Services
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	Service := services.New(srv)
	delivery.New(e, Service)

	e.Logger.Fatal(e.Start(":8000"))
}
