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
	log    logger.Logger
	web    *http.ServeMux
	db     *sql.DB
	respwr respwr.HTTPResponseWriter
}

func NewServer(config config.Config, logger logger.Logger, db *sql.DB) *http.Server {
	s := Server{
		config: config,
		log:    logger,
		web:    http.NewServeMux(),
		db:     db,
		respwr: respwr.NewHTTPResponseWriter(),
	}
	s.setupHandlers()

	server := &http.Server{
		Addr:    config.Port,
		Handler: s.web,
	}
	return server
}
