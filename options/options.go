package options

import (
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

type ServerOption func(*Server)

func WithAddr(addr string) ServerOption {
	return func(s *Server) {
		s.Addr = addr
	}
}

func WithHandler(handler http.Handler) ServerOption {
	return func(s *Server) {
		s.Handler = handler
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.ReadTimeout = timeout
		s.WriteTimeout = timeout
	}
}

// Usage:
//
//	NewServer(
//		WithAddr(":8080"),
//		WithHandler(http.DefaultServeMux),
//		WithTimeout(30*time.Second)
//	)
func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		Server: &http.Server{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
