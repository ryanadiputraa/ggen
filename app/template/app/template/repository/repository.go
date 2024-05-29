package repository

import (
	"github.com/ryanadiputraa/ggen/app/template/app/template"
	"github.com/ryanadiputraa/ggen/app/template/internal/database"
)

type repository struct {
	db database.Service
}

func NewRepository(db database.Service) template.TemplateRepository {
	return &repository{
		db: db,
	}
}
