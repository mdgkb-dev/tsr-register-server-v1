package representative

import "mdgkb/tsr-tegister-server-v1/models"

func (s *Service) Create(item *models.Representative) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() ([]*models.Representative, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.Representative, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Representative) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
