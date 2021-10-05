package history

import "mdgkb/tsr-tegister-server-v1/models"

func (s *Service) Get(id *string) (*models.History, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Create(item *models.History) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	return nil
}
