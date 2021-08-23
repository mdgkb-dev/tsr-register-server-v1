package representative

import (
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Representative) error {
	return s.repository.create(item)
}

func (s *Service) GetAll(pagination *httpHelper.Pagination) ([]*models.Representative, error) {
	if pagination != nil {
		return s.repository.getAll(pagination)
	}
	return s.repository.getOnlyNames()
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
