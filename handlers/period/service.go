package period

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Period) error {
	if item == nil {
		return nil
	}
	return s.repository.create(item)
}

func (s *Service) Update(item *models.Period) error {
	if item == nil {
		return nil
	}
	return s.repository.update(item)
}

func (s *Service) CreateMany(items []*models.Period) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items []*models.Period) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}
