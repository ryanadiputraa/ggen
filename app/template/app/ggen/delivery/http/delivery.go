package http

import (
	"encoding/json"
	"net/http"

	"github.com/ryanadiputraa/ggen/app/template/app/ggen"
)

type httpDelivery struct {
	service ggen.GgenService
}

func NewHTTPDelivery(web *http.ServeMux, service ggen.GgenService) {
	d := &httpDelivery{service: service}

	web.HandleFunc("GET /test", d.TestHandler())
}

func (d *httpDelivery) TestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ggen, err := d.service.GetGgen()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		json, err := json.Marshal(ggen)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write(json)
	}
}
