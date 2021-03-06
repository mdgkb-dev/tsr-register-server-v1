package fileInfo

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/models"

	"github.com/google/uuid"
)

func (s *Service) Create(item *models.FileInfo) error {
	if item == nil {
		return nil
	}
	return s.repository.create(item)
}

func (s *Service) Get(id *string) (*models.FileInfo, error) {
	return s.repository.get(id)
}

func (s *Service) Update(item *models.FileInfo) error {
	if item == nil {
		return nil
	}
	return s.repository.update(item)
}

func (s *Service) Upsert(item *models.FileInfo) error {
	if item == nil {
		return nil
	}
	fmt.Println(item, item)
	return s.repository.upsert(item)
}

func (s *Service) CreateMany(items []*models.FileInfo) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.createMany(items)
}

func (s *Service) UpsertMany(items []*models.FileInfo) error {
	if len(items) == 0 {
		return nil
	}
	return s.repository.upsertMany(items)
}

func (s *Service) Delete(id uuid.NullUUID) error {
	return s.repository.delete(id)
}
