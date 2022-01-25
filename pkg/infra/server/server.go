package server

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Config() (e *echo.Echo) {
	e = echo.New()

	// Timeout middleware is used to timeout
	e.Use(echoMiddleware.Recover())

	return e
}