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
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"

    "%v/configs"
)

type Server struct {
	config *configs.Config
	web    *http.ServeMux
}

func NewHTTPServer(config *configs.Config) *Server {
	return &Server{
		config: config,
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
        if err := server.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    signal.Notify(quit, os.Kill)

    sig := <-quit
    log.Println("received terminate, graceful shutdown ", sig)

    tc, shutdown := context.WithTimeout(context.Background(), 30*time.Second)
    defer shutdown()

    return server.Shutdown(tc)
}`, mod)
}
