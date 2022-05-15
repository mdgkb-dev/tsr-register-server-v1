package register

import (
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/handlers/registerDiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/registerGroup"

	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Register) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = registerGroup.CreateService(s.repository.getDB(), s.helper).UpsertMany(item.RegisterGroups)
	if err != nil {
		return err
	}
	err = registerDiagnosis.CreateService(s.repository.getDB(), s.helper).CreateMany(item.RegisterDiagnosis)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll(userID uuid.UUID) ([]*models.Register, error) {
	items, err := s.repository.getAll(userID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(id string) (*models.Register, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Register) error {
	err := s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	registerGroupService := registerGroup.CreateService(s.repository.getDB(), s.helper)
	err = registerGroupService.UpsertMany(item.RegisterGroups)
	if err != nil {
		return err
	}
	err = registerGroupService.DeleteMany(item.RegisterGroupsForDelete)
	if err != nil {
		return err
	}

	registerDiagnosisService := registerDiagnosis.CreateService(s.repository.getDB(), s.helper)
	err = registerDiagnosisService.UpsertMany(item.RegisterDiagnosis)
	if err != nil {
		return err
	}
	err = registerDiagnosisService.DeleteMany(item.RegisterDiagnosisForDelete)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetValueTypes() (models.ValueTypes, error) {
	items, err := s.repository.getValueTypes()
	if err != nil {
		return nil, err
	}
	return items, nil
}
