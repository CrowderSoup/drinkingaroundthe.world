package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func initIndexHandlerGroup(e *echo.Echo, path string) {
	group := e.Group(path)

	group.GET("", getIndex)
}

func getIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{})
}
