package server

import (
	"net/http"

	templateHTTPDelivery "github.com/ryanadiputraa/ggen/v2/app/template/app/template/delivery/http"
	_templateRepository "github.com/ryanadiputraa/ggen/v2/app/template/app/template/repository"
	_templateService "github.com/ryanadiputraa/ggen/v2/app/template/app/template/service"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/middleware"
)

func (s *Server) setupHandler() http.Handler {
	templateRepository := _templateRepository.NewRepository(s.db)
	templateService := _templateService.NewService(s.logger, templateRepository)
	templateHTTPDelivery.NewHTTPDelivery(s.web, s.respwr, templateService)

	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)
	handler = http.TimeoutHandler(handler, requestTimeoutDuration, "request timeout")
	return handler
}
