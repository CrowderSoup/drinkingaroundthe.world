package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexHandlerGroup struct {
	Group *echo.Group
}

func NewIndexHandlerGroup(e *echo.Echo) *IndexHandlerGroup {
	group := e.Group("")

	group.GET("", getIndex)

	return &IndexHandlerGroup{
		Group: group,
	}
}

func getIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", echo.Map{})
}
