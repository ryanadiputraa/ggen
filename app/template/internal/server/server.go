package server

import (
	"net/http"
	"time"

	"github.com/ryanadiputraa/ggen/v2/app/template/config"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/database"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/logger"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/respwr"
)

const requestTimeoutDuration = time.Second * 30

type Server struct {
	config *config.Config
	logger logger.Logger
	web    *http.ServeMux
	db     database.Service
	respwr respwr.HTTPResponseWriter
}

func NewServer(config *config.Config, logger logger.Logger, db database.Service) *Server {
	return &Server{
		config: config,
		logger: logger,
		web:    http.NewServeMux(),
		db:     db,
		respwr: respwr.NewHTTPResponseWriter(),
	}
}

func (s *Server) ListenAndServe() error {
	handler := s.setupHandler()

	server := &http.Server{
		Addr:    s.config.Port,
		Handler: handler,
	}
	return server.ListenAndServe()
}
