package server

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/ryanadiputraa/ggen/app/template/config"
	"github.com/ryanadiputraa/ggen/app/template/pkg/logger"
	"github.com/ryanadiputraa/ggen/app/template/pkg/respwr"
)

const requestTimeoutDuration = time.Second * 30

type Server struct {
	config *config.Config
	logger logger.Logger
	web    *http.ServeMux
	db     *sql.DB
	respwr respwr.ResponseWriter
}

func NewServer(config *config.Config, db *sql.DB) *Server {
	return &Server{
		config: config,
		logger: logger.NewLogger(time.UTC, os.Stderr),
		web:    http.NewServeMux(),
		db:     db,
		respwr: respwr.NewResponseWriter(),
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
