package patient

import (
	"fmt"
	"mdgkb/tsr-tegister-server-v1/handlers/disability"
	"mdgkb/tsr-tegister-server-v1/handlers/heightWeight"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/patientDiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertySetToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerPropertyToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/registerToPatient"
	"mdgkb/tsr-tegister-server-v1/handlers/representativeToPatient"
	"mdgkb/tsr-tegister-server-v1/helpers/httpHelper"
	"mdgkb/tsr-tegister-server-v1/models"
)

func (s *Service) Create(item *models.Patient) error {
	err := human.CreateService(s.repository.getDB()).Create(item.Human)
	if err != nil {
		return err
	}
	item.HumanID = item.Human.ID
	err = s.repository.create(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()
	err = representativeToPatient.CreateService(s.repository.getDB()).CreateMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = heightWeight.CreateService(s.repository.getDB()).CreateMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = disability.CreateService(s.repository.getDB()).CreateMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = patientDiagnosis.CreateService(s.repository.getDB()).CreateMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = registerToPatient.CreateService(s.repository.getDB()).CreateMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) GetAll(pagination *httpHelper.Pagination) ([]*models.Patient, error) {
	if pagination != nil {
		return s.repository.getAll(pagination)
	}
	return s.repository.getOnlyNames()
}

func (s *Service) Get(id *string) (*models.Patient, error) {
	item, err := s.repository.get(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (s *Service) Update(item *models.Patient) error {
	err := human.CreateService(s.repository.getDB()).Update(item.Human)
	if err != nil {
		return err
	}
	fmt.Println(1)
	item.HumanID = item.Human.ID
	fmt.Println(item)
	err = s.repository.update(item)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(3)
	item.SetIdForChildren()
	fmt.Println(4)

	representativeToPatientService := representativeToPatient.CreateService(s.repository.getDB())
	fmt.Println(5)
	err = representativeToPatientService.UpsertMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	heightWeightService := heightWeight.CreateService(s.repository.getDB())
	err = heightWeightService.UpsertMany(item.HeightWeight)
	if err != nil {
		return err
	}
	err = heightWeightService.DeleteMany(item.HeightWeightForDelete)
	if err != nil {
		return err
	}
	disabilityService := disability.CreateService(s.repository.getDB())
	err = disabilityService.UpsertMany(item.Disabilities)
	if err != nil {
		return err
	}
	err = disabilityService.DeleteMany(item.DisabilitiesForDelete)
	if err != nil {
		return err
	}

	patientDiagnosisService := patientDiagnosis.CreateService(s.repository.getDB())
	err = patientDiagnosisService.UpsertMany(item.PatientDiagnosis)
	if err != nil {
		return err
	}
	err = patientDiagnosisService.DeleteMany(item.PatientDiagnosisForDelete)
	if err != nil {
		return err
	}
	registerToPatientService := registerToPatient.CreateService(s.repository.getDB())
	err = registerToPatientService.UpsertMany(item.RegisterToPatient)
	if err != nil {
		return err
	}
	err = registerToPatientService.DeleteMany(item.RegisterToPatientForDelete)
	if err != nil {
		return err
	}
	err = registerPropertyToPatient.CreateService(s.repository.getDB()).UpsertMany(item.RegisterPropertyToPatient)
	if err != nil {
		return err
	}
	registerPropertySetToPatientService := registerPropertySetToPatient.CreateService(s.repository.getDB())
	err = registerPropertySetToPatientService.UpsertMany(item.RegisterPropertySetToPatient)
	if err != nil {
		return err
	}
	err = registerPropertySetToPatientService.DeleteMany(item.RegisterPropertySetToPatientForDelete)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}

func (s *Service) GetBySearch(query *string) ([]*models.Patient, error) {
	items, err := s.repository.getBySearch(query)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) GetDisabilities() ([]*models.Patient, error) {
	items, err := s.repository.getDisabilities()
	if err != nil {
		return nil, err
	}
	return items, nil
}
