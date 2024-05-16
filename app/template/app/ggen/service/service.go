package service

import "github.com/ryanadiputraa/ggen/app/template/app/ggen"

type service struct {
	repository ggen.GgenRepository
}

func NewService(repository ggen.GgenRepository) ggen.GgenService {
	return &service{
		repository: repository,
	}
}

func (s *service) GetGgen() (val ggen.Ggen, err error) {
	id, err := s.repository.FindByID()
	if err != nil {
		return
	}

	return ggen.NewGgeen(id), nil
}
