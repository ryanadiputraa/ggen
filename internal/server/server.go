package server

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/internal/server", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}
	return writer.WriteToFile(template(mod), path, "server.go")
}

func template(mod string) string {
	return fmt.Sprintf(`package server

import (
    "context"
    "net/http"
    "os"
    "os/signal"
    "time"

    "%[1]v/configs"
    "%[1]v/pkg/logger"
)

type Server struct {
	config *configs.Config
    log    logger.Logger
	web    *http.ServeMux
}

func NewHTTPServer(config *configs.Config, log logger.Logger) *Server {
	return &Server{
		config: config,
        log:    log,
		web:    http.NewServeMux(),
	}
}

func (s *Server) ServeHTTP() error {
	server := &http.Server{
		Addr:         s.config.Port,
		Handler:      s.web,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

    go func() {
        s.log.Info("starting server on port ", s.config.Port)
        if err := server.ListenAndServe(); err != nil {
            s.log.Fatal(err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    signal.Notify(quit, os.Kill)

    sig := <-quit
    s.log.Info("received terminate, graceful shutdown ", sig)

    tc, shutdown := context.WithTimeout(context.Background(), 30*time.Second)
    defer shutdown()

    return server.Shutdown(tc)
}`, mod)
}
