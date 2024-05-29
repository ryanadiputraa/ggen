package main

import (
	"os"
	"time"

	"github.com/ryanadiputraa/ggen/app/template/config"
	"github.com/ryanadiputraa/ggen/app/template/internal/database"
	"github.com/ryanadiputraa/ggen/app/template/internal/logger"
	"github.com/ryanadiputraa/ggen/app/template/internal/server"
)

func main() {
	log := logger.New(time.UTC, os.Stderr)

	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(c)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(c, log, db)
	log.Info("server running on port %v", c.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
