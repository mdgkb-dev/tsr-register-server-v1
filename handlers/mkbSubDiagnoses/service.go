package mkbSubDiagnoses

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetAll() (models.MkbSubDiagnoses, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.MkbSubDiagnosis, error) {
	return s.repository.get(id)
}
