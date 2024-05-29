package http

import (
	"net/http"

	"github.com/ryanadiputraa/ggen/app/template/app/template"
	"github.com/ryanadiputraa/ggen/app/template/internal/respwr"
)

type httpDelivery struct {
	respwr  respwr.HTTPResponseWriter
	service template.TemplateService
}

func NewHTTPDelivery(web *http.ServeMux, respwr respwr.HTTPResponseWriter, service template.TemplateService) {
	d := &httpDelivery{
		respwr:  respwr,
		service: service,
	}

	web.HandleFunc("GET /test", d.TestHandler())
}

func (d *httpDelivery) TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d.respwr.WriteResponseData(w, http.StatusOK, nil)
	}
}
