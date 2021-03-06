package registerGroup

import (
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/handlers/registerProperty"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.RegisterGroup) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
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

	//registerPropertyToRegisterGroupService := registerPropertyToRegisterGroup.CreateService(s.helper)
	//err = registerPropertyToRegisterGroupService.UpsertMany(item.RegisterPropertyToRegisterGroup)
	//if err != nil {
	//	return err
	//}
	//err = registerPropertyToRegisterGroupService.DeleteMany(item.RegisterPropertyToRegisterGroupForDelete)
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.RegisterGroups) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	registerPropertyService := registerProperty.CreateService(s.helper)
	err = registerPropertyService.UpsertMany(items.GetRegisterProperties())
	if err != nil {
		return err
	}
	err = registerPropertyService.DeleteMany(items.GetRegisterPropertiesForDelete())
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteMany(idPool []uuid.UUID) error {
	if len(idPool) == 0 {
		return nil
	}
	return s.repository.deleteMany(idPool)
}
