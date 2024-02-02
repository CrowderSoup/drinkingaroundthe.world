package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/CrowderSoup/drinkingaroundthe.world/web/handlers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer(e *echo.Echo) *Server {
	e.Renderer = echoview.New(goview.Config{
		Root:      "web/views",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"now": time.Now,
		},
	})

	e.Use(middleware.Logger())

	e.Static("/static", "web/static")

	// Init handlers
	handlers.InitializeHandlers(e)

	return &Server{
		echo: e,
	}
}

func (s *Server) Start(port string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start the Server
	go func() {
		err := s.echo.Start(fmt.Sprintf(":%s", port))
		if err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.echo.Logger.Info("attempting graceful shutdown...")

	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
	}
}
