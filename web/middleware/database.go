package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DatabaseMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := c.(*DrinksContext)

			cc.GetRecord("session")

			return nil
		}
	}
}
