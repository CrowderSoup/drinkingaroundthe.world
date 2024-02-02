package middleware

import "github.com/labstack/echo/v4"

type DrinksContext struct {
	echo.Context
	records map[string]any
}

func DrinksContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &DrinksContext{
			c,
			make(map[string]any),
		}
		return next(cc)
	}
}

func (dc *DrinksContext) AddRecord(key string, value any) {
	dc.records[key] = value
}

func (dc *DrinksContext) GetRecord(key string) (any, bool) {
	value, ok := dc.records[key]

	return value, ok
}
