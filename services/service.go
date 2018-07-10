package services

import "tkp/modules/sibur"

type Service interface {
	GetTKP(inPath string) (interface{}, error)
}

type service struct {
	siburModule sibur.API
}

func New(sm sibur.API) Service {
	return &service{
		siburModule: sm,
	}
}

func (s *service) GetTKP(inPath string) (interface{}, error) {
	return s.siburModule.GetTKP(inPath)
}
