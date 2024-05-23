package http

import (
	"net/http"

	"github.com/ryanadiputraa/ggen/app/template/app/ggen"
	"github.com/ryanadiputraa/ggen/app/template/pkg/respwr"
)

type httpDelivery struct {
	respwr  respwr.ResponseWriter
	service ggen.GgenService
}

func NewHTTPDelivery(web *http.ServeMux, respwr respwr.ResponseWriter, service ggen.GgenService) {
	d := &httpDelivery{
		respwr:  respwr,
		service: service,
	}

	web.HandleFunc("GET /test", d.TestHandler())
}

func (d *httpDelivery) TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ggen, err := d.service.GetGgen(r.Context())
		if err != nil {
			d.respwr.WriteErrMessage(w, http.StatusInternalServerError, "internal server error")
			return
		}

		d.respwr.WriteResponseData(w, http.StatusOK, ggen)
	}
}
