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
	return writer.WriteToFile(template(mod), path, "main.go")
}

func template(mod string) string {
	return fmt.Sprintf(`package main

import (
    "%[1]v/configs"
    "%[1]v/internal/server"
    "%[1]v/pkg/logger"
    "%[1]v/pkg/db/postgres"
)

func main() {
	log := logger.NewLogger()

	config, err := configs.LoadConfig("yml", "configs/config.yml")
	if err != nil {
		log.Fatal("load config: ", err)
	}

    db, err := postgres.NewDB(config)
	if err != nil {
		log.Fatal("db connection: ", err)
	}

	server := server.NewHTTPServer(config, log, db)
	if err := server.ServeHTTP(); err != nil {
		log.Fatal("start server: ", err)
	}
}`, mod)
}
