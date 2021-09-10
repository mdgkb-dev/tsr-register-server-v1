package register

import (
	"mdgkb/tsr-tegister-server-v1/handlers/registerDiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/registerGroupToRegister"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Register) error {
	err := s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = registerGroupToRegister.CreateService(s.repository.getDB()).CreateMany(item.RegisterGroupToRegister)
	if err != nil {
		return err
	}
	err = registerDiagnosis.CreateService(s.repository.getDB()).CreateMany(item.RegisterDiagnosis)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() ([]*models.Register, error) {
	items, err := s.repository.getAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) Get(queryFilter *httpHelper.QueryFilter) (*models.Register, error) {
	item, err := s.repository.get(queryFilter)
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

	registerPropertyToRegisterGroupService := registerGroupToRegister.CreateService(s.repository.getDB())
	err = registerPropertyToRegisterGroupService.UpsertMany(item.RegisterGroupToRegister)
	if err != nil {
		return err
	}
	err = registerPropertyToRegisterGroupService.DeleteMany(item.RegisterGroupToRegisterForDelete)
	if err != nil {
		return err
	}

	registerDiagnosisService := registerDiagnosis.CreateService(s.repository.getDB())
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
