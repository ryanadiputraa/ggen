package server

import (
	"net/http"

	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/middleware"
)

func (s *Server) setupHandlers() http.Handler {
	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)
	handler = http.TimeoutHandler(handler, requestTimeoutDuration, "request timeout")
	return handler
}
