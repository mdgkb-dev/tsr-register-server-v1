package mkbDiagnoses

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetAll() (models.MkbDiagnoses, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.MkbDiagnosis, error) {
	return s.repository.get(id)
}
