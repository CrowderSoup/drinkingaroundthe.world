package middleware

import (
	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*DrinksContext)

		session, err := services.GetSession("drinks", cc)
		if err != nil {
			return err
		}

		cc.Set("session", session)

		return next(cc)
	}
}
