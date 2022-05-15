package representative

import (
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/representativeToPatient"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Representative) error {
	err := human.CreateService(s.repository.getDB(), s.helper).Create(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = representativeToPatient.CreateService(s.repository.getDB(), s.helper).CreateMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll() (models.RepresentativesWithCount, error) {
	return s.repository.getAll()
}

func (s *Service) GetOnlyNames() (models.RepresentativesWithCount, error) {
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
	err := human.CreateService(s.repository.getDB(), s.helper).Update(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	representativeToPatientService := representativeToPatient.CreateService(s.repository.getDB(), s.helper)
	err = representativeToPatientService.UpsertMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetBySearch(query *string) ([]*models.Representative, error) {
	queryRu := s.helper.Util.TranslitToRu(*query)
	items, err := s.repository.getBySearch(&queryRu)
	if err != nil {
		return nil, err
	}
	return items, nil
}
