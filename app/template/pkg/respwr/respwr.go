package respwr

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct{}

type ResponseWriter interface {
	WriteResponseData(w http.ResponseWriter, code int, data any)
	WriteErrMessage(w http.ResponseWriter, code int, message string)
	WriteErrDetails(w http.ResponseWriter, code int, message string, errMap map[string]string)
}

type ResponseData struct {
	Data any `json:"data"`
}

type ErrMessage struct {
	Message string `json:"message"`
}

type ErrDetails struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func NewResponseWriter() ResponseWriter {
	return &responseWriter{}
}

func (rw *responseWriter) setHeader(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func (rw *responseWriter) WriteResponseData(w http.ResponseWriter, code int, data any) {
	rw.setHeader(w, code)
	json.NewEncoder(w).Encode(ResponseData{
		Data: data,
	})
}

func (rw *responseWriter) WriteErrMessage(w http.ResponseWriter, code int, message string) {
	rw.setHeader(w, code)
	json.NewEncoder(w).Encode(ErrMessage{
		Message: message,
	})
}

func (rw *responseWriter) WriteErrDetails(w http.ResponseWriter, code int, message string, errMap map[string]string) {
	rw.setHeader(w, code)
	json.NewEncoder(w).Encode(ErrDetails{
		Message: message,
		Errors:  errMap,
	})
}
