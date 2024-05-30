package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ryanadiputraa/ggen/v2/app/template/config"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/logger"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

const requestTimeoutDuration = time.Second * 30

type Server struct {
	config config.Config
	logger logger.Logger
	web    *http.ServeMux
	db     *sql.DB
	respwr respwr.HTTPResponseWriter
}

func NewServer(config config.Config, logger logger.Logger, db *sql.DB) *Server {
	return &Server{
		config: config,
		logger: logger,
		web:    http.NewServeMux(),
		db:     db,
		respwr: respwr.NewHTTPResponseWriter(),
	}
}

func (s *Server) ListenAndServe() (err error) {
	handler := s.setupHandlers()

	server := &http.Server{
		Addr:    s.config.Port,
		Handler: handler,
	}

	return server.ListenAndServe()
}
