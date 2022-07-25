package users

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registersusers"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.User) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = registersusers.CreateService(s.helper).CreateMany(item.RegistersUsers)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.Users, error) {
	return s.repository.getAll()
}

func (s *Service) Get(id *string) (*models.User, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.User) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	registersUsersService := registersusers.CreateService(s.helper)
	err = registersUsersService.UpsertMany(item.RegistersUsers)
	if err != nil {
		return err
	}
	err = registersUsersService.DeleteMany(item.RegistersUsersForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
