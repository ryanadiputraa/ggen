package server

import (
	ggenHTTPDelivery "github.com/ryanadiputraa/ggen/app/template/app/ggen/delivery/http"
	_ggenRepository "github.com/ryanadiputraa/ggen/app/template/app/ggen/repository"
	_ggenService "github.com/ryanadiputraa/ggen/app/template/app/ggen/service"
)

func (s *Server) setHandlers() {
	ggenRepository := _ggenRepository.NewRepository(s.db)
	ggenService := _ggenService.NewService(ggenRepository)
	ggenHTTPDelivery.NewHTTPDelivery(s.web, ggenService)
}
