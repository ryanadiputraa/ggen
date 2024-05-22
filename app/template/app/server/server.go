package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ryanadiputraa/ggen/app/template/app/middleware"
	"github.com/ryanadiputraa/ggen/app/template/config"
)

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
	s.setHandlers()
	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)

	server := &http.Server{
		Addr:         s.config.Port,
		Handler:      handler,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}
	return server.ListenAndServe()
}
