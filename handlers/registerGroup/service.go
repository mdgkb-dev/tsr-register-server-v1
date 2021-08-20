package registerGroup

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToRegisterGroup"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.RegisterGroup) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = registerPropertyToRegisterGroup.CreateService(s.repository.getDB()).CreateMany(item.RegisterPropertyToRegisterGroup)
	return err
}

func (s *Service) GetAll() ([]*models.RegisterGroup, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.RegisterGroup, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.RegisterGroup) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	registerPropertyToRegisterGroupService := registerPropertyToRegisterGroup.CreateService(s.repository.getDB())
	err = registerPropertyToRegisterGroupService.UpsertMany(item.RegisterPropertyToRegisterGroup)
	if err != nil {
		return err
	}
	err = registerPropertyToRegisterGroupService.DeleteMany(item.RegisterPropertyToRegisterGroupForDelete)
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
