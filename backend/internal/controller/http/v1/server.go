package v1

import (
	"fmt"
	"true-kw/config"
)

type Server struct {
	useCases *UseCases
	config   config.AppConfig
}

func NewServer(useCases *UseCases, config config.AppConfig) *Server {
	return &Server{
		useCases: useCases,
		config:   config,
	}
}

func (s *Server) Serve() {
	r := NewRouter(s.useCases)
	r.Run(fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
}
