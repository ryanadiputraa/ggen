package cmd

import (
	"fmt"

	"github.com/ryanadiputraa/ggen/internal/writer"
)

func Write(mod, name string) error {
	path := fmt.Sprintf("%v/cmd/api", name)
	if err := writer.CreateDirectory(path); err != nil {
		return err
	}
	return writer.WriteToFile(template(mod), path)
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
