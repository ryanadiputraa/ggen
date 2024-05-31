package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ryanadiputraa/ggen/v2/app/template/config"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/logger"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

const requestTimeoutDuration = time.Second * 30

type Server struct {
	config config.Config
	log    logger.Logger
	web    *http.ServeMux
	db     *sql.DB
	respwr respwr.HTTPResponseWriter
}

func NewServer(config config.Config, logger logger.Logger, db *sql.DB) *Server {
	return &Server{
		config: config,
		log:    logger,
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

	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		s.log.Info("start shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if err := server.Shutdown(ctx); err == context.DeadlineExceeded {
			s.log.Fatal("server shutdown; context exeeded")
		}
	}()

	s.log.Info("starting server on port", s.config.Port)
	return server.ListenAndServe()
}
