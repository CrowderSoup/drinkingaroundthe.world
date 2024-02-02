package handlers

import "github.com/labstack/echo/v4"

func InitializeHandlers(e *echo.Echo) {
	// Init all handler groups
	initIndexHandlerGroup(e, "")
	initAuthHandlerGroup(e, "/auth")
}
