package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Server struct {
	router *echo.Echo
	port   string
}

type Option func(*Server)

func WithPort(port string) Option {
	return func(server *Server) {
		server.port = port
	}
}

func NewRESTService(router *echo.Echo, opts ...Option) *Server {
	const (
		defaultPort = "5000"
	)

	s := &Server{
		router: router,
		port:   defaultPort,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Run() error {
	return s.router.Start(fmt.Sprintf("127.0.0.1:%s", s.port))
}
