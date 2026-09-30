package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const (
	defaultAddr            = ":80"
	defaultReadTimeout     = 3 * time.Second
	defaultWriteTimeout    = 3 * time.Second
	defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	addr            string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration

	App    *fiber.App
	logger *logrus.Logger
}

func New(logger *logrus.Logger, opts ...Option) *Server {
	s := &Server{
		addr:            defaultAddr,
		readTimeout:     defaultReadTimeout,
		writeTimeout:    defaultWriteTimeout,
		shutdownTimeout: defaultShutdownTimeout,
		App:             nil,
		logger:          logger,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.App = fiber.New(fiber.Config{
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		JSONDecoder:  json.Unmarshal,
		JSONEncoder:  json.Marshal,
	})

	return s
}

func (s *Server) Run() {
	s.logger.Infof("Starting restapi server on %s...\n", s.addr)
	go func() {
		err := s.App.Listen(s.addr)
		s.logger.Error(err, "Error restapi server error")
	}()
	s.logger.Infof("Started")
}

func (s *Server) Stop() {
	s.logger.Infof("Stopping restapi server\n")
	err := s.App.ShutdownWithTimeout(s.shutdownTimeout)
	if err != nil && !errors.Is(err, context.Canceled) {
		s.logger.Error(err, "Stop restapi server error: ShutdownWithTimeout")
	}

	s.logger.Info("Stopped")
}
