package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ryanadiputraa/ggen/app/template/config"
)

const (
	maxOpenConns    = 60
	connMaxLifeTime = 120
	maxIdleConn     = 30
	connMaxIdleTime = 20
)

func NewPostgres(config *config.Config) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC",
		config.PostgresHost,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDB,
		config.PostgresPort,
	)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifeTime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxIdleTime(connMaxIdleTime)

	return
}
