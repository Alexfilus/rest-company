package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/rueian/rueidis"
	"rest-company/config"
	"rest-company/pkg/logger"
)

type Server struct {
	app    *fiber.App
	cfg    *config.Config
	redis  rueidis.Client
	logger logger.Logger
}

func New(cfg *config.Config, redis rueidis.Client, logger logger.Logger) *Server {
	return &Server{
		app:    fiber.New(),
		cfg:    cfg,
		redis:  redis,
		logger: logger,
	}
}

func (s *Server) Run() error {
	s.MapHandlers()
	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.app.Listen(s.cfg.Server.Port); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	s.logger.Info("Server Exited Properly")
	return s.app.Shutdown()
}
