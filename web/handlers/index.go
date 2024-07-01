package handlers

import (
	"net/http"

	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/CrowderSoup/drinkingaroundthe.world/web/middleware"
	"github.com/labstack/echo/v4"
)

func initIndexHandlerGroup(e *echo.Echo, path string) {
	group := e.Group(path)

	group.GET("", getIndex)
}

func getIndex(c echo.Context) error {
	drinksContext := c.(*middleware.DrinksContext)
	session := drinksContext.Get("session").(*services.Session)

	loggedIn := session.GetValue("LoggedIn")
	if loggedIn == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/auth")
	}

	return c.Render(http.StatusOK, "index", echo.Map{})
}
