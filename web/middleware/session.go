package middleware

import (
	"github.com/CrowderSoup/drinkingaroundthe.world/services"
	"github.com/labstack/echo/v4"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*DrinksContext)

		session, err := services.GetSession("drinks", cc)
		if err != nil {
			return err
		}

		existingSessionId := session.GetValue("ID")
		if existingSessionId == nil {
			// Set a unique ID for this session
			id, err := gonanoid.New()
			if err != nil {
				return err
			}
			session.SetValue("ID", id, true)
		}

		cc.Set("session", session)

		return next(cc)
	}
}
