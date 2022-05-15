package registerPropertyRadio

import (
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyOthers"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.RegisterPropertyRadio) error {
	return s.repository.create(item)
}

func (s *Service) GetAll() ([]*models.RegisterPropertyRadio, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id *string) (*models.RegisterPropertyRadio, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.RegisterPropertyRadio) error {
	return s.repository.update(item)
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) UpsertMany(items models.RegisterPropertyRadios) error {
	if len(items) == 0 {
		return nil
	}
	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	registerPropertyOthersService := registerPropertyOthers.CreateService(s.repository.getDB(), s.helper)
	err = registerPropertyOthersService.UpsertMany(items.GetRegisterPropertyOthers())
	if err != nil {
		return err
	}
	err = registerPropertyOthersService.DeleteMany(items.GetRegisterPropertyOthersForDelete())
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