package service

import (
	"context"

	"github.com/ryanadiputraa/ggen/app/template/app/ggen"
	"github.com/ryanadiputraa/ggen/app/template/pkg/logger"
)

type service struct {
	logger     logger.Logger
	repository ggen.GgenRepository
}

func NewService(logger logger.Logger, repository ggen.GgenRepository) ggen.GgenService {
	return &service{
		logger:     logger,
		repository: repository,
	}
}

func (s *service) GetGgen(ctx context.Context) (val ggen.Ggen, err error) {
	id, err := s.repository.FindByID(ctx)
	if err != nil {
		return
	}
	s.logger.Info("fetched id: ", id)

	return ggen.NewGgeen(id), nil
}
