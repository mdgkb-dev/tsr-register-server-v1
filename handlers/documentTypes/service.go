package documentTypes

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.DocumentType) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() ([]*models.DocumentType, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.DocumentType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.DocumentType) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
