package server

import (
	"net/http"

	ggenHTTPDelivery "github.com/ryanadiputraa/ggen/app/template/app/ggen/delivery/http"
	_ggenRepository "github.com/ryanadiputraa/ggen/app/template/app/ggen/repository"
	_ggenService "github.com/ryanadiputraa/ggen/app/template/app/ggen/service"
	"github.com/ryanadiputraa/ggen/app/template/app/middleware"
)

func (s *Server) setupHandler() http.Handler {
	ggenRepository := _ggenRepository.NewRepository(s.db)
	ggenService := _ggenService.NewService(ggenRepository)
	ggenHTTPDelivery.NewHTTPDelivery(s.web, ggenService)

	handler := middleware.CORSMiddleware(s.web)
	handler = middleware.ThrottleMiddleware(handler)
	handler = http.TimeoutHandler(handler, requestTimeoutDuration, "request timeout")
	return handler
}
