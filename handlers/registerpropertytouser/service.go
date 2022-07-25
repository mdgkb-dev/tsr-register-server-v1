package registerpropertytouser

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.RegisterPropertyToUser) error {
	return s.repository.create(item)
}

func (s *Service) Delete(item *models.RegisterPropertyToUser) error {
	return s.repository.delete(item)
}
