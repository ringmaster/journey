package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoute(e *echo.Echo) {

	e.Use(middleware.CORS())

	root := &RootHandler{}
	e.Add("GET", "/", root.Get)

}
