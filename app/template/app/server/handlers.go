package server

import (
	_healthcheckHandler "github.com/ryanadiputraa/ggen/v2/app/template/app/healthcheck/handler"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/middleware"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

func (s *Server) setupHandlers() {
	respwr := respwr.NewHTTPResponseWriter()
	middlewares := middleware.NewMiddlewares()

	healthcheckHandler := _healthcheckHandler.NewHTTPHandler(respwr)

	s.web.Handle("GET /healthcheck", middlewares(healthcheckHandler.Healthcheck()))
}
