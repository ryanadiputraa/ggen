package main

import (
	"log"

	"github.com/ryanadiputraa/ggen/app/template/app/server"
	"github.com/ryanadiputraa/ggen/app/template/config"
	"github.com/ryanadiputraa/ggen/app/template/pkg/db"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := db.NewPostgres(c)
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(c, postgres)
	log.Printf("server running on port %v", c.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
