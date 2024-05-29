package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/ryanadiputraa/ggen/app/template/config"
)

type Service interface {
	// Close terminate database connection
	// It return an error if connection cannot be closed
	Close() error
}

type service struct {
	db *sql.DB
}

const (
	maxOpenConns    = 60
	connMaxLifeTime = 120
	maxIdleConn     = 30
	connMaxIdleTime = 20
)

var (
	instance *service
)

func New(config *config.Config) (s Service, err error) {
	if instance != nil {
		return instance, nil
	}
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresHost,
		config.Port,
		config.PostgresDB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	// Health check by pinging database with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err = db.PingContext(ctx); err != nil {
		return
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	s = &service{
		db: db,
	}
	return
}

// Close close database connection
// If an error occurs when closing the connection, it will return the error
func (s *service) Close() error {
	return nil
}
