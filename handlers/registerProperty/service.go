package registerProperty

import "mdgkb/tsr-tegister-server-v1/models"

func (s *Service) Create(item *models.RegisterProperty) error {
	return s.repository.create(item)
}

func (s *Service) GetAll(registerId *string) ([]*models.RegisterProperty, error) {
	items, err := s.repository.getAll(registerId)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.RegisterProperty, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.RegisterProperty) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetValueTypes() ([]*models.ValueType, error) {
	items, err := s.repository.getValueTypes()
	if err != nil {
		return nil, err
	}
	return items, nil
}
