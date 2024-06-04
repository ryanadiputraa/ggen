package server

import (
	"net/http"

	_healthcheckHandler "github.com/ryanadiputraa/ggen/v2/app/template/app/healthcheck/handler"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/middleware"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

func setupHandlers(s Server) http.Handler {
	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)
	handler = http.TimeoutHandler(handler, requestTimeoutDuration, "request timeout")

	respwr := respwr.NewHTTPResponseWriter()
	healthcheckHandler := _healthcheckHandler.NewHTTPHandler(respwr)

	s.web.Handle("GET /healthcheck", healthcheckHandler.Healthcheck())
	return handler
}
