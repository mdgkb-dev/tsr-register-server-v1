package mkbgroups

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) GetAll() (models.MkbGroups, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id string) (*models.MkbGroup, error) {
	return s.repository.get(id)
}
