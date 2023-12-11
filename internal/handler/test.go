package handler

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/internal/test/handler", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}

	return writer.WriteToFile(template(mod), path, "handler.go")
}

func template(mod string) string {
	return fmt.Sprintf(`package handler

import (
    "encoding/json"
    "net/http"

    "%[1]v/internal/domain/test"
)

type handler struct{}

func NewHandler(web *http.ServeMux) {
    h := &handler{}

    web.HandleFunc("/api/test", h.TestHandler)
}

func (h *handler) TestHandler(w http.ResponseWriter, r *http.Request) {
    t := test.Test{
        Message: "Hello World!",
    }

    v, err := json.Marshal(t)
    if err != nil {
        w.Write([]byte("internal server error"))
        return
    }

    w.Write(v)
    return
}`, mod)
}
