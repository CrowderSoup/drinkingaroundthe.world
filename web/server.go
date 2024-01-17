package web

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) Start(port string) error {
	return s.echo.Start(fmt.Sprintf(":%s", port))
}
