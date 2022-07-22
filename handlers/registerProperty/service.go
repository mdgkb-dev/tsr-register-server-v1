package registerProperty

import (
	"github.com/google/uuid"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyExamples"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyMeasures"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyRadio"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertySet"
	"mdgkb/tsr-tegister-server-v1/models"
)

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

func (s *Service) UpsertMany(items models.RegisterProperties) error {
	if len(items) == 0 {
		return nil
	}

	err := s.repository.upsertMany(items)
	if err != nil {
		return err
	}
	items.SetIdForChildren()
	registerPropertyRadioService := registerPropertyRadio.CreateService(s.helper)
	err = registerPropertyRadioService.UpsertMany(items.GetRegisterPropertyRadios())
	if err != nil {
		return err
	}
	err = registerPropertyRadioService.DeleteMany(items.GetRegisterPropertyRadioForDelete())
	if err != nil {
		return err
	}
	registerPropertySetService := registerPropertySet.CreateService(s.helper)
	err = registerPropertySetService.UpsertMany(items.GetRegisterPropertySets())
	if err != nil {
		return err
	}
	err = registerPropertySetService.DeleteMany(items.GetRegisterPropertySetForDelete())
	if err != nil {
		return err
	}

	registerPropertyExamplesService := registerPropertyExamples.CreateService(s.helper)
	err = registerPropertyExamplesService.UpsertMany(items.GetRegisterPropertyExamples())
	if err != nil {
		return err
	}
	err = registerPropertyExamplesService.DeleteMany(items.GetRegisterPropertyExamplesForDelete())
	if err != nil {
		return err
	}

	registerPropertyMeasuresService := registerPropertyMeasures.CreateService(s.helper)
	err = registerPropertyMeasuresService.UpsertMany(items.GetRegisterPropertyMeasures())
	if err != nil {
		return err
	}
	err = registerPropertyMeasuresService.DeleteMany(items.GetRegisterPropertyMeasuresForDelete())
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
