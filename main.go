package main

import (
	"github.com/CrowderSoup/drinkingaroundthe.world/cmd"
	"github.com/CrowderSoup/drinkingaroundthe.world/web"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(cmd.Module, web.Module, fx.NopLogger)
	app.Run()
}
