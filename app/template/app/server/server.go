package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ryanadiputraa/ggen/app/template/config"
)

const requestTimeoutDuration = time.Second * 30

type Server struct {
	config *config.Config
	web    *http.ServeMux
	db     *sql.DB
}

func NewServer(config *config.Config, db *sql.DB) *Server {
	return &Server{
		config: config,
		web:    http.NewServeMux(),
		db:     db,
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
