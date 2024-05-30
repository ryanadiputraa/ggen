package main

import (
	"os"
	"time"

	"github.com/ryanadiputraa/ggen/v2/app/template/app/server"
	"github.com/ryanadiputraa/ggen/v2/app/template/config"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/db"
	"github.com/ryanadiputraa/ggen/v2/app/template/pkg/logger"
)

func main() {
	log := logger.New(time.UTC, os.Stderr)

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal("load config:", err)
	}

	db, err := db.NewPostgres()
	if err != nil {
		log.Fatal("postgres:", err)
	}

	s := server.NewServer(c, log, db)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("start server:", err)
	}
}
