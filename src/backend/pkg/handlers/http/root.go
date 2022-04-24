package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RootHandler struct {
}

func (c *RootHandler) Get(e echo.Context) error {
	id := e.QueryParam("id")
	return e.JSON(http.StatusOK, id)
}
