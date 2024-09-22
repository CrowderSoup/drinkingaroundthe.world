package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitializeHandlers(e *echo.Echo, db *gorm.DB) {
	// Init all handler groups
	initIndexHandlerGroup(e, "")
	initAuthHandlerGroup(e, db, "/auth")
}
