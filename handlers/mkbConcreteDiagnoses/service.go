package mkbConcreteDiagnoses

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetAll() (models.MkbConcreteDiagnoses, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.MkbConcreteDiagnosis, error) {
	return s.repository.get(id)
}
