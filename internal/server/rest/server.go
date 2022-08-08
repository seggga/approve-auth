package rest

import (
	"context"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

// Server ...
type Server struct {
	srv    http.Server
	logger *zap.SugaredLogger
}

// NewServer ..
func NewServer(h http.Handler, logger *zap.SugaredLogger) *Server {
	s := &Server{
		logger: logger,
	}

	s.srv = http.Server{
		Addr:              ":" + os.Getenv("AUTH_PORT_3000_TCP_PORT"),
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
	return s
}

// Start ...
func (s *Server) Start() {
	if s.srv.Addr == ":" {
		s.logger.Fatal("variable AUTH_PORT_3000_TCP_PORT not set")
	}

	go s.srv.ListenAndServe()
	s.logger.Infof("REST service started on port %s", s.srv.Addr)
}

// Stop ...
func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	s.srv.Shutdown(ctx)
	cancel()
	s.logger.Info("REST server stopped")
}
