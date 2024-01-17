package web

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewServer,
		echo.New,
	),
)
