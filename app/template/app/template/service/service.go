package service

import (
	"github.com/ryanadiputraa/ggen/v2/app/template/app/template"
	"github.com/ryanadiputraa/ggen/v2/app/template/internal/logger"
)

type service struct {
	logger     logger.Logger
	repository template.TemplateRepository
}

func NewService(logger logger.Logger, repository template.TemplateRepository) template.TemplateService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}
