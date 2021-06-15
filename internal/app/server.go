package app

import (
	"context"
	"fmt"
	"net/http"
	"wiltonassessment/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.ConfigMap, mux *http.ServeMux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%s", cfg.HTTPData.Host, cfg.HTTPData.Port),
			Handler:      mux,
			ReadTimeout:  cfg.HTTPData.ReadTimeout,
			WriteTimeout: cfg.HTTPData.WriteTimeout,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
