package repository

import (
	"github.com/ryanadiputraa/ggen/v2/app/template/app/template"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/database"
)

type repository struct {
	db database.Service
}

func NewRepository(db database.Service) template.TemplateRepository {
	return &repository{
		db: db,
	}
}
