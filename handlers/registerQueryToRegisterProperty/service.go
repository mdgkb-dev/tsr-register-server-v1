package registerQueryToRegisterProperty

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) CreateMany(items []*models.RegisterQueryToRegisterProperty) error {
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items []*models.RegisterQueryToRegisterProperty) error {
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []string) error {
	if len(idPool) == 0 {
		return nil
	}

	return s.repository.deleteMany(idPool)
}
