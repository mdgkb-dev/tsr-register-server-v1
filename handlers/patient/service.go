package patient

import (
	"mdgkb/tsr-tegister-server-v1/handlers/anthropometryData"
	"mdgkb/tsr-tegister-server-v1/handlers/disability"
	"mdgkb/tsr-tegister-server-v1/handlers/human"
	"mdgkb/tsr-tegister-server-v1/handlers/patientDiagnosis"
	"mdgkb/tsr-tegister-server-v1/handlers/representativeToPatient"
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
	err = anthropometryData.CreateService(s.repository.getDB()).CreateMany(item.AnthropometryData)
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
	return err
}

func (s *Service) GetAll(offset *int) ([]*models.Patient, error) {
	items, err := s.repository.getAll(offset)
	if err != nil {
		return nil, err
	}
	return items, nil
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
	item.HumanID = item.Human.ID
	err = s.repository.update(item)
	if err != nil {
		return err
	}
	item.SetIdForChildren()

	representativeToPatientService := representativeToPatient.CreateService(s.repository.getDB())
	err = representativeToPatientService.UpsertMany(item.RepresentativeToPatient)
	if err != nil {
		return err
	}
	err = representativeToPatientService.DeleteMany(item.RepresentativeToPatientForDelete)
	if err != nil {
		return err
	}
	anthropometryDataService := anthropometryData.CreateService(s.repository.getDB())
	err = anthropometryDataService.UpsertMany(item.AnthropometryData)
	if err != nil {
		return err
	}
	err = anthropometryDataService.DeleteMany(item.AnthropometryDataForDelete)
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
	return nil
}

func (s *Service) Delete(id *string) error {
	return s.repository.delete(id)
}
