package representatives

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/humans"
	"mdgkb/tsr-tegister-server-v1/handlers/representativesdomains"
	"mdgkb/tsr-tegister-server-v1/handlers/representativetopatient"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(c context.Context, item *models.Representative) error {
	err := humans.S.Create(c, item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.Create(c, item)
	if err != nil {
		return err
	}
	err = representativesdomains.S.AddToDomain(c, item.ID)
	if err != nil {
		return err
	}

	return err
}

func (s *Service) GetAll(c context.Context) (models.RepresentativesWithCount, error) {
	return s.repository.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Representative, error) {
	item, err := s.repository.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(c context.Context, item *models.Representative) error {
	err := s.repository.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	representativeToPatientService := representativetopatient.CreateService(s.helper)
	err = representativeToPatientService.UpsertMany(item.PatientsRepresentatives)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.PatientsRepresentativesForDelete)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	return s.repository.Delete(c, id)
}

func (s *Service) GetBySnilsNumber(c context.Context, snils string) (*models.Representative, bool, error) {
	item, err := s.repository.GetBySnilsNumber(c, snils)
	if err != nil {
		return nil, false, err
	}
	exists, err := representativesdomains.S.RepresentativeInDomain(c, item.ID.UUID.String())
	if err != nil {
		return nil, false, err
	}
	return item, exists, nil
}
