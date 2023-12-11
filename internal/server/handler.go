package server

import (
	"fmt"
)

func handlerTemplate(mod string) string {
	return fmt.Sprintf(`package server

import _testHandler "%[1]v/internal/test/handler"

func (s *Server) setupHandlers() {
    _testHandler.NewHandler(s.web)
}`, mod)
}
