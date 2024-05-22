package repository

import (
	"context"
	"database/sql"

	"github.com/ryanadiputraa/ggen/app/template/app/ggen"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) ggen.GgenRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindByID(ctx context.Context) (id string, err error) {
	rows, err := r.db.Query("SELECT 'ggen-id'")
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return "", err
		}
	}

	return
}
