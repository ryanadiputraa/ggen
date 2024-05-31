package http

import (
	"net/http"

	"github.com/ryanadiputraa/ggen/v2/app/template/app/health"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/respwr"
)

type handler struct {
	respwr respwr.HTTPResponseWriter
}

func NewHTTPHandler(respwr respwr.HTTPResponseWriter) handler {
	return handler{
		respwr: respwr,
	}
}

func (h *handler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	data := health.Health{
		Status: "ok",
	}
	h.respwr.WriteResponseData(w, http.StatusOK, data)
}
