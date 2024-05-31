package server

import (
	"net/http"

	_healthHandler "github.com/ryanadiputraa/ggen/v2/app/template/app/health/delivery/http"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/middleware"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

func (s *Server) setupHandlers() http.Handler {
	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)
	handler = http.TimeoutHandler(handler, requestTimeoutDuration, "request timeout")

	respwr := respwr.NewHTTPResponseWriter()
	healthHandler := _healthHandler.NewHTTPHandler(respwr)

	s.web.HandleFunc("GET /test", healthHandler.Healthcheck)
	return handler
}
