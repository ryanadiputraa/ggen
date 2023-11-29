package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/cmd/api", name)

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	write := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' > %s/main.go", template(mod), path))
	return write.Run()
}

func template(mod string) string {
	return fmt.Sprintf(`package main

import (
    "log"

    "%v/internal/server"
)

func main() {
    server := server.NewHTTPServer()
    if err := server.ServeHTTP(); err != nil {
        log.Fatal(err)
    }
}`, mod)
}
