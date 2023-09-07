package patients

import (
	"context"
	"mdgkb/tsr-tegister-server-v1/handlers/humans"
	"mdgkb/tsr-tegister-server-v1/handlers/patientsdomains"
	"mdgkb/tsr-tegister-server-v1/middleware"

	"mdgkb/tsr-tegister-server-v1/handlers/patientdrugregimen"
	"mdgkb/tsr-tegister-server-v1/handlers/representativetopatient"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(c context.Context, item *models.Patient) error {
	err := humans.S.Create(c, item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.Create(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()
	err = representativetopatient.CreateService(s.helper).CreateMany(item.PatientsRepresentatives)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll(c context.Context) (models.PatientsWithCount, error) {
	return s.repository.GetAll(c)
}

func (s *Service) Get(c context.Context, id string) (*models.Patient, error) {
	item, err := s.repository.Get(c, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) GetBySnilsNumber(c context.Context, snils string) (*models.Patient, bool, error) {
	item, err := s.repository.GetBySnilsNumber(c, snils)
	if err != nil {
		return nil, false, err
	}
	exists, err := patientsdomains.S.PatientInDomain(c, item.ID.UUID.String(), middleware.ClaimDomainIDS.FromContext(c))
	if err != nil {
		return nil, false, err
	}
	return item, exists, nil
}

func (s *Service) Update(c context.Context, item *models.Patient) error {
	err := humans.S.Update(c, item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.Update(c, item)
	if err != nil {
		return err
	}
	item.SetIDForChildren()

	representativeToPatientService := representativetopatient.CreateService(s.helper)
	err = representativeToPatientService.UpsertMany(item.PatientsRepresentatives)
	if err != nil {
		return err
	}

	patientDrugRegimenService := patientdrugregimen.CreateService(s.helper)
	err = patientDrugRegimenService.UpsertMany(item.PatientDrugRegimen)
	if err != nil {
		return err
	}
	err = patientDrugRegimenService.DeleteMany(item.PatientDrugRegimenForDelete)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(c context.Context, id string) error {
	err := s.repository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
