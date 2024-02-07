package middleware

import "github.com/labstack/echo/v4"

type DrinksContext struct {
	echo.Context
	records map[string]interface{}
}

func DrinksContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &DrinksContext{
			c,
			make(map[string]interface{}),
		}
		return next(cc)
	}
}

func (dc *DrinksContext) AddRecord(key string, value interface{}) {
	dc.records[key] = value
}

func (dc *DrinksContext) GetRecord(key string) (interface{}, bool) {
	value, ok := dc.records[key]

	return value, ok
}
