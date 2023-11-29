package server

import (
	"fmt"
	"os"
	"os/exec"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/internal/server", name)

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	write := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/server.go", template(mod), path))
	return write.Run()
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
)

type Server struct {
    web *http.ServeMux
}

func NewHTTPServer() *Server {
    return &Server{
        web: http.NewServeMux(),
    }
}

func (s *Server) ServeHTTP() error {
    server := &http.Server{
        Addr:         ":8080",
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
}`)
}
