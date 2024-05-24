package api

import (
	"context"
	"net/http"

	"github.com/SeaRoll/mini-auth/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Server interface {
	Run(ctx context.Context, stop chan bool)
}

type server struct {
	config config.Config
	e      *echo.Echo
}

func NewServer(config config.Config) Server {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	return &server{
		config: config,
		e:      e,
	}
}

func (s *server) Run(ctx context.Context, stop chan bool) {
	go func() {
		if err := s.e.Start(":" + s.config.Port); err != nil && err != http.ErrServerClosed {
			s.e.Logger.Fatal("shutting down the server")
		}
	}()
	<-stop
	if err := s.e.Shutdown(ctx); err != nil {
		s.e.Logger.Fatal(err)
	}
}
