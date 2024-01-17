package web

import (
	"fmt"

	"github.com/CrowderSoup/drinkingaroundthe.world/web/handlers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(e *echo.Echo, indexHandlerGroup *handlers.IndexHandlerGroup) *Server {
	e.Renderer = echoview.New(goview.Config{
		Root:      "web/views",
		Extension: ".html",
		Master:    "layouts/master",
	})

	e.Use(middleware.Logger())

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(port string) error {
	return s.echo.Start(fmt.Sprintf(":%s", port))
}
