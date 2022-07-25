package representativetypes

import (
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.RepresentativeType) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() ([]*models.RepresentativeType, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.RepresentativeType, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.RepresentativeType) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
