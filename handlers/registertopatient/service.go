package registertopatient

import (
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) CreateMany(items []*models.ResearchResult) error {
	if len(items) == 0 {
		return nil
	}

	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items []*models.ResearchResult) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}

func (s *Service) Delete(id string) error {
	return s.repository.delete(id)
}
